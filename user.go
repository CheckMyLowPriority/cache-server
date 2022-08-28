package main

func RegisterNewUser(login, pass string) *User {
	return &User{Login: login, Pass: pass}
}

/*func HandleUser(w http.ResponseWriter, r *http.Request) {
	loginform := r.FormValue("login")
	passform := r.FormValue("pass")
	RegisterNewUser(loginform, passform)
	m, _ := json.Marshal(r)
	json.Unmarshal(m, User{})
}*/
/*func (u *User) AddUserInDb(login, pass string) userId {
	var userId int
	_ = `insert into users value id, login, pass=$1,$2,$3` //id auto, <- login pass into
	_ = `select from users value login, pass where id =$1`
	lastId := LastInsertId
	userId = lastId
	return userId
}*/
/*func (u *User) CheckUserInDb(login, pass string) {
	userid := u.AddUserInDb(login, pass) //there id user was last added in db
	if login == "" && pass == "" {
	}
}*/
