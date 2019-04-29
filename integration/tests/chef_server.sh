#shellcheck disable=SC2034
#shellcheck disable=SC2039
#shellcheck disable=SC2154
#shellcheck disable=SC1091

test_name="chef-server"

source .studio/chef-server-collection

do_create_config() {
    do_create_config_default
    # TODO(jaym) Figure out below
    # Github seems displeased with us constantly asking for the latest liveness agent version
    # For now, I'm hardcoding it
    #log_info "Determining latest release of automate-livenes-agent"
    #latest=$(curl "https://api.github.com/repos/chef/automate-liveness-agent/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    latest="v0.7.7"
    log_info "Downloading automate-livenes-agent/$latest"
    curl -L -o "/tmp/required_recipe.rb" \
        "https://github.com/chef/automate-liveness-agent/releases/download/$latest/automate-liveness-recipe.rb"
    chmod a+r /tmp/required_recipe.rb

    cat << EOF >> "$test_config_path"
[erchef.v1.sys.memory]
max_bytes = 2000000000

[cs_nginx.v1.sys.required_recipe]
enabled = true
path = "/tmp/required_recipe.rb"
EOF
}

do_deploy() {
    chef-automate deploy config.toml \
        --hartifacts "$test_hartifacts_path" \
        --override-origin "$HAB_ORIGIN" \
        --manifest-dir "$test_manifest_path" \
        --enable-chef-server \
        --admin-password chefautomate \
        --accept-terms-and-mlsa
}

do_test_deploy() {
    PATH="/hab/bin:/bin" chef-server-ctl test
    test_chef_server_ctl
    test_knife
    test_cookbook_caching
    converge_chef_client

    # Converging the chef client should run the required recipe which sets
    # up the liveness agent and starts it.
    local PIDFILE="/var/run/automate-liveness-agent.pid"
    local COUNTER=1
    while [[ ! -f "$PIDFILE" ]]; do
        if [[ $COUNTER -ge 30 ]]; then
            log_error "liveness agent pidfile never appeared. Dumping ps and log"
            ps fuax
            cat /var/log/chef/automate-liveness-agent/automate-liveness-agent.log
            return 1
        fi
        echo "Waiting for $PIDFILE to appear $COUNTER/30";
        (( COUNTER++ ))
        sleep 1
    done

    pid=$(cat $PIDFILE)
    if ! ps -p "$pid"
    then
        log_error "liveness agent (pid=$pid) was not found. Dumping ps..."
        ps faux
        return 1
    fi

    do_test_deploy_default
}
