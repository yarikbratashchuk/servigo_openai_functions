** Methods of Servigo SDK

| Servio endpoint | OpenAI API Function Signaturees | Comments |
|--|--|--|
| GetCompanyInfo | GetCompanyInfoFunc |  |
| GetRooms | GetRoomsFunc |  |
| AddRoomReservation | AddRoomReservationFunc |  |
| SetReservationType |  | Not needed on this moment |
| SetReservationBill |  | Not needed on this moment |
| GetAccountConfirm | GetAccountCofirmFunc |  |
| GetRoomTypesList | GetRoomTypesListFunc |  |
| GetGroupRooms | GetGroupRoomsFunc |  |
| AddGroupRoomReservation | AddGroupRoomReservationFunc |  |
| GetLoyaltyCardNumberInfo |  | Not needed on this moment (method is not activated on server) |
| GetReservationInfo | GetReservationInfoFunc |  |
| GetRoomType | GetRoomTypeFunc |  |
| GetClientDocuments |  | Not needed on this moment |
| SetClientDocuments |  | Not needed on this moment (method is not activated on server) |
| CancelReservation | CancelReservationFunc |  |
| GetRoomTypesMinPrice | GetRoomTypesMinPrice |  |
| GetMinPrices | GetMinPricesFunc |  |


------------------------------------------
|        Servigo OpenAI Functions        |
------------------------------------------
```
[
    {
        "name": "GetCompanyInfoFunc",
        "description": "This function retrieves detailed information about a company as the company's Name, Email, Fax, Phone Number and whether it requires prepayment. It is used for querying and accessing specific company data within the system.",
        "parameters": null
    },
    {
        "name": "GetRoomsFunc",
        "description": "The function gets the list of room categories with the number of available rooms of each category according to the input parameters.",
        "parameters": {
            "type": "object",
            "properties": {
                "Adults": {
                    "type": "integer",
                    "description": "Number of adults. Must be equal or greater than 1."
                },
                "ChildAges": {
                    "type": "array",
                    "description": "Array of children's ages",
                    "items": {
                        "type": "integer"
                    }
                },
                "DateArrival": {
                    "type": "string",
                    "description": "Date of arrival in format YYYY-MM-DD"
                },
                "DateDeparture": {
                    "type": "string",
                    "description": "Date of departure in format YYYY-MM-DD"
                },
                "IsExtraBedUsed": {
                    "type": "boolean",
                    "description": "Is there a need for an extra bed."
                },
                "TimeArrival": {
                    "type": "string",
                    "description": "Time of arrival in format HH:MM"
                },
                "TimeDeparture": {
                    "type": "string",
                    "description": "Time of arrival in format HH:MM"
                }
            },
            "required": [
                "DateArrival",
                "DateDeparture",
                "Adults"
            ]
        }
    },
    {
        "name": "GetRoomTypesMinPrice",
        "description": "The function retrieves the minimum possible price by room category for 30 days in advance.",
        "parameters": {
            "type": "object",
            "properties": {
                "RoomTypeIDs": {
                    "type": "array",
                    "description": "Array of room type IDs.",
                    "items": {
                        "type": "integer"
                    }
                }
            },
            "required": [
                "RoomTypeIDs"
            ]
        }
    },
    {
        "name": "GetRoomTypesListFunc",
        "description": "The function returns a list of room categories by hotel.",
        "parameters": null
    },
    {
        "name": "GetRoomTypeFunc",
        "description": "The function returns a full description of the room by ID and localizations",
        "parameters": {
            "type": "object",
            "properties": {
                "RoomTypeID": {
                    "type": "integer",
                    "description": "Room type ID."
                }
            },
            "required": [
                "RoomTypeID"
            ]
        }
    },
    {
        "name": "GetReservationInfoFunc",
        "description": "The function returns information on the created reservation, data is returned only for confirmed reservations.",
        "parameters": {
            "type": "object",
            "properties": {
                "Account": {
                    "type": "string",
                    "description": "Booking ID"
                }
            },
            "required": [
                "Account"
            ]
        }
    },
    {
        "name": "GetMinPricesFunc",
        "description": "The function returns the minimum possible prices by room category for the specified period.",
        "parameters": {
            "type": "object",
            "properties": {
                "EndDate": {
                    "type": "string",
                    "description": "End date of the search period in format YYYY-MM-DD"
                },
                "StartDate": {
                    "type": "string",
                    "description": "Start date of the search period in format YYYY-MM-DD"
                }
            },
            "required": [
                "StartDate",
                "EndDate"
            ]
        }
    },
    {
        "name": "GetGroupRoomsFunc",
        "description": "The function gets the list of room categories with the number of available rooms of each category according to the input parameters.",
        "parameters": {
            "type": "object",
            "properties": {
                "BookingRequests": {
                    "type": "array",
                    "description": "Array of booking requests.",
                    "required": [
                        "DateArrival",
                        "DateDeparture",
                        "Adults"
                    ],
                    "items": {
                        "type": "object",
                        "properties": {
                            "Adults": {
                                "type": "integer",
                                "description": "Number of adults. Must be equal or greater than 1."
                            },
                            "ChildAges": {
                                "type": "array",
                                "description": "Array of children's ages",
                                "items": {
                                    "type": "integer"
                                }
                            },
                            "DateArrival": {
                                "type": "string",
                                "description": "Date of arrival in format YYYY-MM-DD"
                            },
                            "DateDeparture": {
                                "type": "string",
                                "description": "Date of departure in format YYYY-MM-DD"
                            },
                            "IsExtraBedUsed": {
                                "type": "boolean",
                                "description": "Is there a need for an extra bed."
                            },
                            "TimeArrival": {
                                "type": "string",
                                "description": "Time of arrival in format HH:MM"
                            },
                            "TimeDeparture": {
                                "type": "string",
                                "description": "Time of arrival in format HH:MM"
                            }
                        }
                    }
                }
            },
            "required": [
                "BookingRequests"
            ]
        }
    },
    {
        "name": "CancelReservationFunc",
        "description": "The function will cancel (cancel the reservation) of a guest, group or event, if possible.",
        "parameters": {
            "type": "object",
            "properties": {
                "Account": {
                    "type": "string",
                    "description": "Booking ID"
                }
            },
            "required": [
                "Account"
            ]
        }
    },
    {
        "name": "GetAccountCofirmFunc",
        "description": "The function generates and returns a reservation confirmation by generating a compressed zip document or folder in the desired format.",
        "parameters": {
            "type": "object",
            "properties": {
                "Account": {
                    "type": "string",
                    "description": "Booking ID"
                },
                "Format": {
                    "type": "string",
                    "description": "File format. Must be '0' for PDF, '1' - HTML, '2' - Excel, '3' - Doc",
                    "enum": [
                        "0",
                        "1",
                        "2",
                        "3"
                    ]
                }
            },
            "required": [
                "Account",
                "Format"
            ]
        }
    },
    {
        "name": "AddRoomReservationFunc",
        "description": "The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.",
        "parameters": {
            "type": "object",
            "properties": {
                "Address": {
                    "type": "string",
                    "description": "Guest Address"
                },
                "Adults": {
                    "type": "integer",
                    "description": "Number of adults. Must be equal or greater than 1."
                },
                "AgencyCategory": {
                    "type": "string",
                    "description": "The guest category. 0 - not specified (the default category configured for the hotel will be used), 1 - Resident, 2 - Non-resident.",
                    "enum": [
                        "0",
                        "1",
                        "2"
                    ]
                },
                "BirthDate": {
                    "type": "string",
                    "description": "Guest BirthDate in format YYYY-MM-DD"
                },
                "ChildAges": {
                    "type": "array",
                    "description": "Array of children's ages",
                    "items": {
                        "type": "integer"
                    }
                },
                "Comment": {
                    "type": "string",
                    "description": "Any Comment"
                },
                "ContactName": {
                    "type": "string",
                    "description": "Contact Person Name"
                },
                "Country": {
                    "type": "string",
                    "description": "Guest Country"
                },
                "DateArrival": {
                    "type": "string",
                    "description": "Date of arrival in format YYYY-MM-DD"
                },
                "DateDeparture": {
                    "type": "string",
                    "description": "Date of departure in format YYYY-MM-DD"
                },
                "Email": {
                    "type": "string",
                    "description": "Guest Email"
                },
                "Fax": {
                    "type": "string",
                    "description": "Guest Fax"
                },
                "GuestFirstName": {
                    "type": "string",
                    "description": "Guest First Name"
                },
                "GuestLastName": {
                    "type": "string",
                    "description": "Guest Last Name"
                },
                "IsExtraBedUsed": {
                    "type": "boolean",
                    "description": "Is there a need for an extra bed."
                },
                "IsTouristTax": {
                    "type": "string",
                    "description": "A sign of inclusion of the tourist tax in the automatically created invoice. 1 if the tourist tax is included (the guest is traveling for tourist purposes), 0 if the tourist tax is not included (the guest is traveling on business, is on a business trip).",
                    "enum": [
                        "0",
                        "1"
                    ]
                },
                "NeedTransport": {
                    "type": "string",
                    "description": "0 if the guest does not need to order transportation, 1 if the guest needs to order transportation",
                    "enum": [
                        "0",
                        "1"
                    ]
                },
                "PaidType": {
                    "type": "string",
                    "description": "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
                    "enum": [
                        "100",
                        "200",
                        "300"
                    ]
                },
                "Phone": {
                    "type": "string",
                    "description": "Guest Phone"
                },
                "RoomTypeID": {
                    "type": "integer",
                    "description": "Room Type ID"
                },
                "TimeArrival": {
                    "type": "string",
                    "description": "Time of arrival in format HH:MM"
                },
                "TimeDeparture": {
                    "type": "string",
                    "description": "Time of arrival in format HH:MM"
                }
            },
            "required": [
                "DateArrival",
                "DateDeparture",
                "Adults",
                "GuestLastName",
                "GuestFirstName",
                "RoomTypeID",
                "PaidType",
                "NeedTransport",
                "IsTouristTax"
            ]
        }
    },
    {
        "name": "AddGroupRoomReservationFunc",
        "description": "The function creates a new reservation in the database by incoming parameters, and returns the reservation ID.",
        "parameters": {
            "type": "object",
            "properties": {
                "Comment": {
                    "type": "string",
                    "description": "Any Comment"
                },
                "ContactEmail": {
                    "type": "string",
                    "description": "Contact Person Email"
                },
                "ContactInfo": {
                    "type": "string",
                    "description": "Contact Person Info"
                },
                "ContactName": {
                    "type": "string",
                    "description": "Contact Person Name"
                },
                "Country": {
                    "type": "string",
                    "description": "Guest Country"
                },
                "GroupName": {
                    "type": "string",
                    "description": "Group Name"
                },
                "PaidType": {
                    "type": "string",
                    "description": "Payment Type. 100 for cash, 200 for credit card, 300 for cashless",
                    "enum": [
                        "100",
                        "200",
                        "300"
                    ]
                },
                "RoomReservations": {
                    "type": "array",
                    "description": "Array of group reservations",
                    "required": [
                        "DateArrival",
                        "DateDeparture",
                        "Adults",
                        "GuestName",
                        "RoomTypeID",
                        "NeedTransport",
                        "IsTouristTax"
                    ],
                    "items": {
                        "type": "object",
                        "properties": {
                            "Adults": {
                                "type": "integer",
                                "description": "Number of adults. Must be equal or greater than 1."
                            },
                            "AgencyCategory": {
                                "type": "string",
                                "description": "The guest category. 0 - not specified (the default category configured for the hotel will be used), 1 - Resident, 2 - Non-resident.",
                                "enum": [
                                    "0",
                                    "1",
                                    "2"
                                ]
                            },
                            "ChildAges": {
                                "type": "array",
                                "description": "Array of children's ages",
                                "items": {
                                    "type": "integer"
                                }
                            },
                            "DateArrival": {
                                "type": "string",
                                "description": "Date of arrival in format YYYY-MM-DD"
                            },
                            "DateDeparture": {
                                "type": "string",
                                "description": "Date of departure in format YYYY-MM-DD"
                            },
                            "GuestName": {
                                "type": "string",
                                "description": "Guest Full Name"
                            },
                            "IsExtraBedUsed": {
                                "type": "boolean",
                                "description": "Is there a need for an extra bed."
                            },
                            "IsTouristTax": {
                                "type": "string",
                                "description": "A sign of inclusion of the tourist tax in the automatically created invoice. 1 if the tourist tax is included (the guest is traveling for tourist purposes), 0 if the tourist tax is not included (the guest is traveling on business, is on a business trip).",
                                "enum": [
                                    "0",
                                    "1"
                                ]
                            },
                            "NeedTransport": {
                                "type": "string",
                                "description": "0 if the guest does not need to order transportation, 1 if the guest needs to order transportation",
                                "enum": [
                                    "0",
                                    "1"
                                ]
                            },
                            "RoomTypeID": {
                                "type": "integer",
                                "description": "Room Type ID"
                            },
                            "TimeArrival": {
                                "type": "string",
                                "description": "Time of arrival in format HH:MM"
                            },
                            "TimeDeparture": {
                                "type": "string",
                                "description": "Time of arrival in format HH:MM"
                            }
                        }
                    }
                }
            },
            "required": [
                "RoomReservations",
                "GroupName",
                "PaidType"
            ]
        }
    }
]
```