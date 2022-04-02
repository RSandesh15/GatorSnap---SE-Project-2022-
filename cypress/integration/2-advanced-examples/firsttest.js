const { cyan } = require("@material-ui/core/colors")

describe('Test 1', function (){
    it('Test 1', function(){
        cy.visit("http://localhost:3000/")
        cy.get('button').should('contain', 'Button')
        cy.title().should('eq','GatorSnaps!')
        
    })
})