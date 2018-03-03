function getUserComplaint(id, complaintType) {
    const user = getUser(id);
    return user.complaints[complaintType];
}