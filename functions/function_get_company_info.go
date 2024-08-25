package functions

import (
	"github.com/sashabaranov/go-openai"
)

//SDK Method is available to return also company's ID, CodeID but it is not important information for customer
func newGetCompanyInfoFunc() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name: "get_company_info",
		Description: `This function retrieves detailed information about a company as the company's Name, Email, Fax, Phone Number and whether it requires prepayment. It is used for querying and accessing specific company data within the system.`,
	}
}

