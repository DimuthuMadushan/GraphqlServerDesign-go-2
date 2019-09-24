package gqlServerType2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graph-gophers/graphql-go"
)

const Schema = `
	schema {
		query: Query
	}
	
	type Query {
		employee(id: ID!):Employee!
	}
	
	type Employee {
		id: ID!
		name: String!
		age: String!
		email: String!
	}

`

type employee struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
}

func (emp employee) ID() graphql.ID {
	return graphql.ID(emp.Id)
}

func (emp employee) NAME() string {
	return emp.Name
}

func (emp employee) AGE() string {
	return emp.Age
}

func (emp employee) EMAIL() string {
	return emp.Email
}

type Resolver struct{}

func (r *Resolver) Employee(ctx context.Context, args struct{ ID string }) (employee, error) {
	employees := employee{}
	id := fmt.Sprint(args.ID)
	link := "http://localhost:9090/employeemgt/employee/" + id
	response, err := http.Get(link)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			responseString := string(content)
			temp := fmt.Sprint(responseString)
			//fmt.Println(temp)
			err = json.Unmarshal([]byte(temp), &employees)
			if err != nil {
				err := fmt.Errorf("user with id=%s does not exist", args.ID)
				return employee{}, err
			} else {
				//fmt.Println(employees)
				return employees, err

			}
		}
	}
	return employee{}, err
}
