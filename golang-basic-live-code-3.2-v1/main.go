package main

import (
	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
	"errors"
	"fmt"
	"strconv"
)

type Learnly interface {
	RegisterUser(user model.User) error
	LoginUser(email, password string) (model.User, error)
	GetLessonsByCategory(email, category string) ([]model.Lesson, error)
}

type learnlyApp struct {
	users   []model.User
	lessons []model.Lesson
}

func NewLearnly(users []model.User, lessons []model.Lesson) learnlyApp {
	return learnlyApp{users, lessons}
}

func (l *learnlyApp) GetUser() model.Users {
	return l.users
}

func (l *learnlyApp) AddLesson(lesson model.Lesson) {
	l.lessons = append(l.lessons, lesson)
}

func (l *learnlyApp) GetLesson() model.Lessons {
	return l.lessons
}

func (l *learnlyApp) Reset() {
	l.users = []model.User{}
	l.lessons = []model.Lesson{}
}

func (l *learnlyApp) Validate(u model.User) error {

	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if u.Age < 0 || u.Age > 120 {
		return errors.New("age should be between 0 and 120")
	}
	if u.Gender != "Male" && u.Gender != "Female" {
		return errors.New("gender should be either Male or Female")
	}
	return nil
}

func (l *learnlyApp) RegisterUser(user model.User) error {
	if err := l.Validate(user); err != nil {
		return err
	}
	for _, u := range l.users {
		if u.Email == user.Email {
			return errors.New("email already registered")
		}
	}
	l.users = append(l.users, user)
	return nil
}

func (l *learnlyApp) LoginUser(email, password string) (model.User, error) {
	users := l.GetUser()
	user := model.User{}
	var err error
	for i := 0; i < len(users); i++ {
		if users[i].Email == email && users[i].Password == password {
			users[i].Session = true
			user = users[i]
		} else {
			err = errors.New("invalid email or password")
		}
	}
	return user, err // TODO: replace this
}

// nope // TODO: replace this

func (l *learnlyApp) GetLessonsByCategory(email, category string) ([]model.Lesson, error) {
	lesson := []model.Lesson{}
	users := l.GetUser()
	err := errors.New("you must login first")
	for _, u := range users {
		if u.Email == email && u.Session == true {
			lessons := l.GetLesson()
			for _, l := range lessons {
				if l.Category == category {
					lesson = append(lesson, l)
				}
			}
			err = nil
			break
		} else {
			return nil, err
		}
	}
	return lesson, err // TODO: replace this
}

func main() {
	app := NewLearnly([]model.User{}, []model.Lesson{})

	var choice int
	for {
		helper.ClearScreen()
		fmt.Println("User: ", app.GetUser())
		fmt.Println("Lesson: ", app.GetLesson())

		fmt.Println("Welcome to LearnlyApp!")
		fmt.Println("1. Register user")
		fmt.Println("2. Login user")
		fmt.Println("3. Add lesson")
		fmt.Println("4. Get lesson by category")
		fmt.Println("5. Exit")
		fmt.Print("Choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			helper.ClearScreen()
			helper.InputScan("1. Register\n")
			name := helper.InputScan("\t- Name: ")
			email := helper.InputScan("\t- Email: ")
			password := helper.InputScan("\t- Password: ")
			age := helper.InputScan("\t- Age: ")
			ageInt, _ := strconv.Atoi(age)
			gender := helper.InputScan("\t- Gender (Male/Female): ")
			user := model.User{
				Name:     name,
				Email:    email,
				Password: password,
				Age:      ageInt,
				Gender:   gender,
				Session:  false,
			}
			err := app.RegisterUser(user)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User registered successfully")
			}
			helper.Delay(5)
		case 2:
			helper.ClearScreen()
			helper.InputScan("2. Login\n")
			email := helper.InputScan("\t- Email: ")
			password := helper.InputScan("\t- Password: ")
			user, err := app.LoginUser(email, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Logged in as %s\n", user.Name)
			}
			helper.Delay(5)
		case 3:
			helper.ClearScreen()
			helper.InputScan("3. Add lesson\n")
			title := helper.InputScan("\t- Title: ")
			description := helper.InputScan("\t- Description: ")
			category := helper.InputScan("\t- Category: ")
			difficulty := helper.InputScan("\t- Difficulty: ")
			difficultyInt, _ := strconv.Atoi(difficulty)

			lesson := model.Lesson{
				Title:       title,
				Description: description,
				Category:    category,
				Difficulty:  difficultyInt,
			}

			app.AddLesson(lesson)
			helper.Delay(5)
		case 4:
			helper.ClearScreen()
			helper.InputScan("4. Get lesson by category\n")
			email := helper.InputScan("\t- Email: ")
			category := helper.InputScan("\t- Category: ")

			res, err := app.GetLessonsByCategory(email, category)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Lesson: ", res)
			}
			helper.Delay(5)
		case 5:
			fmt.Println("Thank you for using LearnlyApp!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
