package service

import (
	"agenda/entity"
	"fmt"
	"log"
	"os"
	"strings"
)

var my_name, my_password string
var Login_flag bool
var All_name []string

var log_file *os.File

func GetFlag() bool {
	return Login_flag
}

func Init() {
	entity.Init()

	logFile, err := os.OpenFile("service/agenda.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	log_file = logFile
	if err != nil {
		log.Fatalln("open file error !")
	}
	//debugLog := log.New(log_file,"[Operation]",log.LstdFlags)
	tmp := entity.LN_ReadFromFile()
	if len(tmp) == 0 {
		Login_flag = false
	} else {
		Login_flag = true
		my_name = strings.Replace(tmp[0], "\n", "", -1)
	}

}

func RegisterUser(name string, password string, email string, phone string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	i := entity.RegisterUser(name, password, email, phone)
	if i {
		debugLog.Println(name, " register successfully!")
	} else {
		debugLog.Println(name, " register failed!")
	}
	defer log_file.Close()
}

func Cancellmeeting(title string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	_, flag, _ := entity.Query_meeting_by_title(title)
	if flag == false {
		fmt.Println(title, "doesn't exists!")
		debugLog.Println(my_name, " cancell meeting ", title, " failed!")
		os.Exit(1)
	} else {
		is_flag := entity.Cancell_meeting(title, my_name)
		if is_flag == false {
			fmt.Println(title, " is not sponsor by you, so you can't cancell the meeting!")
			debugLog.Println(my_name, " cancell meeting ", title, " failed!")
			os.Exit(1)
		}
		debugLog.Println(my_name, " cancell meeting ", title, " successfully!")
		fmt.Println("Cancell meeting successfully!")
	}
	defer log_file.Close()
}

func Addparticipator(name string, title string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	tmp_m := entity.Query_meeting_by_sponsor(my_name)
	if len(tmp_m) != 0 {
		for i := 0; i < len(tmp_m); i++ {
			if entity.GetTitle(tmp_m[i]) == title {
				if entity.Add_participator(name, title) {
					debugLog.Println(my_name, " add participator ", name, " to ", title, " successfully!")
					fmt.Println("Add participator to ", title, " successfully!")
				} else {
					debugLog.Println(my_name, " add participator ", name, " to ", title, " failed!")
					fmt.Println("Time is overlap!") //time overlap
				}
			}
		}
	} else {
		debugLog.Println(my_name, " add participator ", name, " to ", title, " failed!")
		fmt.Println("You didn't initiate the meeting!")
	}
	defer log_file.Close()
}

func Createmeeting(title string, start string, end string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	if !entity.IsValid(entity.StringToDate(start)) || !entity.IsValid(entity.StringToDate(end)) ||
		entity.Date_LessThan(entity.StringToDate(end), entity.StringToDate(start)) {
		debugLog.Println(my_name, " create meeting ", title, " failed!")
		fmt.Println("Time is error!")
	} else {
		if entity.Create_meeting(title, entity.StringToDate(start), entity.StringToDate(end), my_name) {
			debugLog.Println(my_name, " create meeting ", title, " successfully!")
		} else {
			debugLog.Println(my_name, " create meeting ", title, " failed!")
		}
	}
	defer log_file.Close()
}

func Deleteuser() {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	debugLog.Println(my_name, " log off account successfully!")
	entity.Empty_login()
	entity.Delete_user(my_name)
	fmt.Println("log off successfully!")
	defer log_file.Close()
}

func Emptymeeting() {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	i := entity.Empty_meeting(my_name)
	if i == 0 {
		debugLog.Println(my_name, " empty meeting failed!")
		fmt.Println("You didn't sponsor any meeting!")
	} else {
		debugLog.Println(my_name, " empty meeting successfully!")
		fmt.Println("Empty meeting successfully!")
	}
	defer log_file.Close()
}

func Exitmeeting(title string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	a := entity.Exit_meeting(my_name, title)
	if a == 1 {
		debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("It doesn't exists this meeting!")
	} else if a == 2 {
		debugLog.Println(my_name, " exit ", title, " meeting successfully!")
		fmt.Println("Exit meeting successfully!")
	} else if a == 3 {
		debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("You are the sponsor, you can't exit meeting!")
	} else {
		debugLog.Println(my_name, " exit ", title, " meeting failed!")
		fmt.Println("You don't attend the meeting!")
	}
	defer log_file.Close()
}

func Login(name string, password string) {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	tmp_u, flag, _ := entity.Query_user(name)
	if flag == true {
		my_name = name
		my_password = password
		if entity.GetPassword(tmp_u) != password {
			debugLog.Println(name, " log in failed!")
			fmt.Println("Password is wrong!")
		} else {
			debugLog.Println(name, " log in successfully!")
			fmt.Println("Log in successfully!\nWelcome to Agenda!")
		}
	} else {
		debugLog.Println(name, " log in failed!")
		fmt.Println("You don't register!")
	}
	entity.LN_WriteToFile(name)
	defer log_file.Close()
}

func Logout() {
	debugLog := log.New(log_file, "[Operation]", log.LstdFlags)
	debugLog.Println(my_name, " log out successfully!")
	fmt.Println("Log out successfully!")
	entity.Empty_login()
	defer log_file.Close()
}



