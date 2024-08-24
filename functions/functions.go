package functions

import "github.com/sashabaranov/go-openai"

func CreateFunctionList() []openai.FunctionDefinition {
	return []openai.FunctionDefinition{
		newGetCompanyInfoFunc(),
		newGetRooms(),
		getRoomTypesMinPrice(),
		newGetRoomTypesList(),
		newGetRoomType(),
		newGetReservationInfo(),
		newGetMinPrices(),
		newGetGroupRooms(),
		newCancelReservation(),
		newGetAccountConfirm(),
		newAddRoomReservation(),
		newAddGroupRoomReservation(),
	}
}


