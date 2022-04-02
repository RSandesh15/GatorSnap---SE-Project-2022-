const { cyan } = require("@material-ui/core/colors")

describe('Test 1', function (){
    it('Test 1', function(){
        cy.visit("http://localhost:3000/")
        cy.get('button').should('contain', 'Button')
        cy.title().should('eq','GatorSnaps!')
        
    })
})

describe('Test 2', function (){
    it('Test 2', function(){
        cy.visit("http://localhost:3000/SellerLogin")
        cy.get('button').should('contain', 'Button')   
    })
})

describe('Test 3', function (){
    it('Test 3', function(){
        cy.visit("http://localhost:3000/SellerUploadPage")
        cy.get('button').should('contain', 'Button')   
    })
})

describe('Test 4', function (){
    it('Test 4', function(){
        cy.visit("http://localhost:3000/SignUp")
        cy.get('button').should('contain', 'Button')   
    })
})

describe('Test 5', function (){
    it('Test 5', function(){
        cy.visit("http://localhost:3000/Checkout")
        cy.get('button').should('contain', 'Button')   
    })
})

describe('Test 6', function (){
    it('Test 6', function(){
        cy.visit("http://localhost:3000/Dashboard")
        cy.get('button').should('contain', 'Button')   
    })
})

describe('Test 7', function (){
    it('Test 7', function(){
        cy.visit("http://localhost:3000/SellerLogin")
        cy.get('button').should('contain', 'Button')   
    })
})