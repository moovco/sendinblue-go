/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package lib

type PostContactInfoContacts struct {
	Success []string `json:"success,omitempty"`
	Failure []string `json:"failure,omitempty"`
	// Displays the count of total number of contacts removed from list when user opts for \"all\" option.
	Total int64 `json:"total,omitempty"`
	// Id of the process created to remove contacts from list when user opts for \"all\" option.
	ProcessId int64 `json:"processId,omitempty"`
}
