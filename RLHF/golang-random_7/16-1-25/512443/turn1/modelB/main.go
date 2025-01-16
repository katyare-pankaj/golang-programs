package main

import "fmt"

func main() {
	endpoints := []map[string]string{
		{
			"Endpoint": "/users",
			"Method":   "GET",
			"Params":   "user_id: integer (optional)",
			"Response": "user_list: [user object]",
		},
		{
			"Endpoint": "/users/{user_id}",
			"Method":   "GET",
			"Params":   "",
			"Response": "user_object: {user details}",
		},
		{
			"Endpoint": "/users/{user_id}",
			"Method":   "POST",
			"Params":   "user_data: {user object}",
			"Response": "updated_user_object: {updated user details}",
		},
		{
			"Endpoint": "/users/{user_id}",
			"Method":   "DELETE",
			"Params":   "",
			"Response": "message: User deleted successfully",
		},
	}

	apiVersion := "v1"
	fmt.Printf("# API Documentation: Version %s\n", apiVersion)
	fmt.Println("---")

	for _, endpoint := range endpoints {
		endpointString := fmt.Sprintf("### %s %s\n", endpoint["Method"], endpoint["Endpoint"])
		endpointString += fmt.Sprintf("**Parameters:**\n  %s\n\n", endpoint["Params"])
		endpointString += fmt.Sprintf("**Response:**\n  %s\n\n", endpoint["Response"])
		fmt.Println(endpointString)
	}
}
