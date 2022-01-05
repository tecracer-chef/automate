describe('chef server', () => {
  let adminIdToken = '';
  const now = Cypress.moment().format('MMDDYYhhmmss');
  const cypressPrefix = 'infra';
  const serverID = 'chef-manage';
  const serverName = 'chef manage';
  const orgID = 'demoorg';
  const orgName = 'demoorg';
  const serverFQDN = 'https://ec2-18-117-112-129.us-east-2.compute.amazonaws.com';
  const serverIP = '18-117-112-129';
  const adminUser = 'kallol';
  const adminKey = Cypress.env('AUTOMATE_INFRA_ADMIN_KEY').replace(/\\n/g, '\n');
  const tabNames = [
    'Roles',
    'Environments',
    'Data Bags',
    'Clients',
    'Policyfiles',
    'Policy Groups',
    'Cookbooks'];
  const webuiKey = Cypress.env('INFRA_WEBUI_KEY').replace(/\\n/g, '\n');

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
          ip_address: serverIP,
          webui_key: webuiKey
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

  describe('chef organizations details page', () => {
    it('displays org details', () => {
      cy.get('.page-title').contains(orgName);
    });

    it('on page load Cookbook is in active state', () => {
      cy.get('.nav-tab').contains('Cookbooks').should('have.class', 'active');
    });

    tabNames.forEach((val) => {
      it(`can switch to ${val} tab`, () => {
          cy.get('.nav-tab').contains(val).click();
      });
    });

    it('lists of Cookbook', () => {
      cy.get('.cookbooks').then(($cookbook) => {
        if ($cookbook.hasClass('empty-section')) {
          cy.get('[data-cy=cookbooks-table-container]').should('not.be.visible');
            cy.get('.empty-section').should('be.visible');
            cy.get('.empty-section p').contains('No cookbooks available');
        } else {
          cy.get('[data-cy=cookbooks-table-container] chef-th').contains('Name');
          cy.get('[data-cy=cookbooks-table-container] chef-th').contains('Cookbook Version');
        }
      });
    });

  });
});
