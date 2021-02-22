# Api doc


# # To Create a new User

Endpoint:- /api/user 

Request Method :- POST

Content-Type: :- application/json

Payload: 
{
  "firstname":"your name",
    "lastname":"your lastname",
    "password":"password",
    "email":"mymail@gmail.com"
}

Response:
{
		"message": "Account Created"
}
    
   
    
# # Login to the todo app

Endpoint:- /api/login 

Request Method :- POST

Content-Type: :- application/json

Payload: 
{
    "password":"password",
    "email":"mymail@gmail.com"
}

Response:
{
  "redirecturl": "/web/todo.html",
  "id": "123445566"
}




# # Add Todo 

Endpoint:- /api/todo

Request Method :- POST

Content-Type: :- application/json

sample Request: http://loacalhost:port/api/todo?id=1223333 
***id will be provided  once you login***
***todoPriority :- ['HIGH','LOW','MEDIUM'] any of this list

Payload: 
{
"ID": "8abdfe7f6812d92c",
"todoPriority": "HIGH", // HIGH,LOW,MEDIUM
"description": "descp",
"enddate": "2021-12-22 00:00:00",
"title": "New todo",
"todoID": "00232c57bdc05170"
}

Response:
{
 "message": "Todo Created"
}

    
# # Get Todo 


Endpoint:- /api/todo 

Request Method :- GET

sample Request: http://loacalhost:port/api/todo?id=1223333 
***id will be provided  once you login***

Response:
{
"ID": "8abdfe7f6812d92c",
"todoPriority": "HIGH",
"description": "descp",
"enddate": "2021-12-22 00:00:00",
"title": "New todo",
"is_completed": "false",
"todoID": "00232c57bdc05170"
}






    
# # Update Todo 



Endpoint:- /api/todo 
Request Method :- PUT

sample Request: http://loacalhost:port/api/todo?id=1223333 
***id will be provided  once you login***

Request payload:
{
"ID": "8abdfe7f6812d92c",
"todoPriority": "HIGH",
"description": "descp",
"enddate": "2021-12-22 00:00:00",
"title": "New todo",
"todoID": "00232c57bdc05170"
}

Response:
{
 "message": "Todo updated"
}





# # Delete Todo 



Endpoint:- /api/todo 

Request Method :- DELETE

sample Request: http://loacalhost:port/api/todo?id=1223333&todoid=12233344  
***id will be provided  once you login***
***you can get todoid  from get api***


Response:
{
		"message": "todo deleted"
}



# # Todo Completed



Endpoint:- /api/todo/completed

Request Method :- PUT

sample Request: http://loacalhost:port/api/todo/completed?id=1223333&todoid=12233344 
***id will be provided  once you login***
***you can get todoid  from get api***

Response:
{
  "message": "Todo updated to completed"
}

    
    
    
    
