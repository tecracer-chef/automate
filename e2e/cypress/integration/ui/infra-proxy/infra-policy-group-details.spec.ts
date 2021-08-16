describe('infra policy group details', () => {
  let adminIdToken = '';
  const serverID = 'chef-server-dev-test';
  const serverName = 'chef server dev';
  const orgID = 'chef-org-dev';
  const orgName = '4thcoffee';
  const serverFQDN = 'ec2-34-219-25-251.us-west-2.compute.amazonaws.com';
  const serverIP = '34.219.25.251';
  const adminUser = 'chefadmin';
  const adminKey = Cypress.env('AUTOMATE_INFRA_ADMIN_KEY').replace(/\\n/g, '\n');
  let policies: any;
  let policyGroupName = '';
  let policyFilesCount: number;

  before(() => {
    cy.adminLogin('/').then(() => {
      const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
      adminIdToken = admin.id_token;

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: '/api/v0/infra/servers',
        body: {
          id: serverID,
          name: serverName,
          fqdn: serverFQDN,
          ip_address: serverIP
        }
      }).then((resp) => {
        if (resp.status === 200 && resp.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}`,
            body: {
              id: serverID
            }
          });
        }
      });

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: `/api/v0/infra/servers/${serverID}/orgs`,
        body: {
          id: orgID,
          server_id: serverID,
          name: orgName,
          admin_user: adminUser,
          admin_key: adminKey
        }
      }).then((response) => {
        if (response.status === 200 && response.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}`,
            body: {
              id: orgID,
              server_id: serverID
            }
          });
        }
      });
      cy.visit(`/infrastructure/chef-servers/${serverID}/organizations/${orgID}`);
      cy.get('app-welcome-modal').invoke('hide');
    });
    cy.restoreStorage();
  });

  beforeEach(() => {
    cy.restoreStorage();
  });

  afterEach(() => {
    cy.saveStorage();
  });

  function getPolicyFile() {
    return cy.request({
      auth: { bearer: adminIdToken },
      failOnStatusCode: false,
      method: 'GET',
      url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}/policyfiles`
    });
  }

  function checkResponse(response: any) {
    if (response.body.policies.length === 0) {
      cy.get('[data-cy=empty-list]').should('be.visible');
    } else {
      policies = response.body.policies;
      policyGroupName = response.body.policies[0].policy_group;
      cy.get('[data-cy=policy-group-table-container] chef-th').contains('Policy Group');
      cy.get('[data-cy=policy-group-table-container] chef-th').contains('Number of Policyfiles');
      return true;
    }
  }

  function getPolicyGroupDetails(policyName: string) {
    const url = `/api/v0/infra/servers/${serverID}/orgs/${orgID}/policygroups`;
    return cy.request({
      auth: { bearer: adminIdToken },
      failOnStatusCode: false,
      method: 'GET',
      url: `${url}/${policyName}`
    });
  }

  function checkPolicyGroupDetailsResponse(response: any) {
    if (response.body.name !== '') {
      policyFilesCount = response.body.policies.length;
      return true;
    }
  }

  describe('infra policy groups list page', () => {
    it('displays org details', () => {
      cy.get('.page-title').contains(orgName);
    });

    // policy groups tabs specs
    it('can switch to policy groups tab', () => {
      cy.get('.nav-tab').contains('Policy Groups').click();
    });

    it('can check if policy groups has list or not', () => {
        getPolicyFile().then((response) => {
          checkResponse(response);
        });
    });

    it('can go to details page', () => {
      if (policyGroupName !== '') {
        cy.get('[data-cy=policy-group-table-container] chef-td').contains(policyGroupName).click();
      }
    });

    it('displays policyfile details', () => {
      if (policyGroupName !== '') {
        getPolicyGroupDetails(policyGroupName).then((response) => {
          checkPolicyGroupDetailsResponse(response);
        });
        cy.get('.page-title').contains(policyGroupName);
        cy.get('[data-cy=policy-group-server]').contains(serverID);
        cy.get('[data-cy=policy-group-org]').contains(orgID);
      }
    });
  });
});
