import { nodejson, uuidv4 } from '../../../support/helpers';

describe('Node manager service ', () => {

let id: any;
  it('add secret', () => {
    cy.request({
      headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
      method: 'POST',
      url: 'api/v0/secrets',
      body: {
        id: 'test',
        name: 'test',
        type: 'ssh',
        data: [
          {key: 'username', value: 'root'},
          {key: 'password', value: '123456'},
          {key: 'key', value: null}
        ],
        tags: []
      }
    }).then((response) => {
      id = response.body.id;
    });
  });

  let nodeId;

  it('add node', () => {
    cy.request({
      headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
      method: 'POST',
      url: 'api/v0/nodes/bulk-create',
      body: {
        nodes: [{
          name: 'test-0.0.0.0',
          manager: 'automate',
          target_config: {
          backend: 'ssh',
          secrets: [],
          port: 22,
          sudo: false,
          hosts: ['localhost']
          },
          tags: []
        }]
      }
    }).then((response) => {
      nodeId = response.body.ids[0];
    });
  });

  it('upload file', () => {
    // const formData = new FormData();
    // formData.set('file', new File(['data'], 'upload.txt'), 'upload.txt');

    // cy.request({
    //   headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
    //   method: 'POST',
    //   url: 'api/v0/compliance/profiles?contentType=application/x-gzip&owner=admin',
    //   body: {
    //     query: {
    //       filter_map: []
    //     }
    //   }
    // }).then((response) => {
    //   expect(response.body.summary.valid).to.equal(true);
    // });

    cy.uploadFileRequest(
      'testproxy-0.1.0.tar.gz',
      'testproxy-0.1.0.tar.gz',
      'demoFileUploadRequest'
    );

    // cy.wait(30000)

    cy.wait('@demoFileUploadRequest', {timeout: 120000}).then((response) => {
      expect(response.status).to.eq(200);
    });
  });

  it('get manager id', () => {
    cy.request({
      headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
      method: 'POST',
      url: 'api/v0/nodemanagers/search',
      body: {filter_map: [{key: 'manager_type', values: ['automate']}], sort: 'date_added'}
    }).then((response) => {
      id = response.body.managers[0].id;
    });
  });

  it('run scan', () => {
    cy.wait(10000);
    cy.request({
      headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
      method: 'POST',
      url: 'api/v0/compliance/scanner/jobs',
      body: {
          type: 'exec',
          tags: [],
          name: 'test',
          profiles: ['compliance://admin/testProxy#0.1.0'],
          node_selectors: [{
            manager_id: id,
            filters: [{
              key: 'name',
              values: ['test-0.0.0.0'],
              exclude: false
            }]
          }],
          recurrence: ''
      }
    }).then((response) => {
      id = response.body.id;
    });
  });

  it('wait for scan', () => {

    var dt = new Date();

    let year  = dt.getFullYear();
    let month = (dt.getMonth() + 1).toString().padStart(2, "0");
    let day   = dt.getDate().toString().padStart(2, "0");
    let dayPrev   = (dt.getDate() - 1).toString().padStart(2, "0");
    cy.wait(10000);
    cy.request({
      headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
      method: 'POST',
      url: 'api/v0/compliance/reporting/stats/summary',
      body: {
        filters: [{
          type: 'job_id',
          values: [id]},
        {
          type: 'start_time',
          values: [`${year}-${month}-${dayPrev}T00:00:00Z`]}, {type: 'end_time', values: [`${year}-${month}-${day}T23:59:59Z`]}]}
    }).then((response) => {
      cy.log('response.body:::::::', response.body);
      id = response.body.id;
    });
  });
});

declare global {
  interface Window {
    XMLHttpRequest: any;
  }
}

Cypress.Commands.add('uploadFileRequest', (fileToUpload: any, uniqueName: any, aliasName: any) => {
  const data = new FormData();

  cy.server()
    .route({
      method: 'POST',
      url: 'api/v0/compliance/profiles?contentType=application/x-gzip&owner=admin',
    })
    .as(aliasName)
    .window()
    .then((win) => {
      cy.fixture(fileToUpload, 'binary')
        .then((binary) => Cypress.Blob.binaryStringToBlob(binary))
        .then((blob) => {
          const xhr = new win.XMLHttpRequest();

          data.set('file', blob, fileToUpload);
          // var formData = new FormData();
          // formData.append('testproxy-0.1.0.tar.gz', fileToUpload);

          xhr.open('POST', 'api/v0/compliance/profiles?contentType=application/x-gzip&owner=admin');

          xhr.setRequestHeader('api-token', Cypress.env('ADMIN_TOKEN'));
          xhr.timeout = 300000;

          xhr.send(data);
        });
    });
})

