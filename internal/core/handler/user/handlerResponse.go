package user

import "github.com/gin-gonic/gin"

type Response struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

const (
	InvalidData         = "Invalid data."
	ValidationFailed    = "Validation failed."
	FailedToAddUser     = "Failed to add user."
	FailedToGetUsers    = "Failed to get users."
	FailedToUpdateUser  = "Failed to update user."
	UserUpdated         = "User updated."
	FailedToDeleteUser  = "Failed to delete user."
	UserDeleted         = "User deleted."
	FailedToGetUserByID = "Failed to get user by id."
	UserRetrieved       = "User retrieved."
	InvalidID           = "Invalid id."
	InvalidPagination   = "Invalid pagination parameters."
	InvalidFilters      = "Invalid filter parameters."
	UserProfileUpdated  = "User profile updated."
	UserPasswordUpdated = "User password updated."
	UserNotExists       = "User not exists."
)

func sendResponse(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
