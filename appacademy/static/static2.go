package static

// START OMIT
type User struct {
	Name       string
	Age        int
	Complaints map[string]Complaint
}

type Complaint struct {
	Type    string
	Memo    string
	DaysOld int
}

func getUserComplaint(id string, complaintType string) Complaint {
	// END OMIT
	var user = getUser(id)
	return user.Complaints[complaintType]
}

func getUser(id string) User {
	return User{}
}
