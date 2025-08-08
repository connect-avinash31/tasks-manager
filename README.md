# tasks-manager
1. This repo contains the code related to assignment of task-manager
2. This is simple application which support the CRUD operation on Task resource and store in DB.
3. It's using postgress to store the data, so need to install postgress as backend 

4. Operations supported 
	1. Add of the task using API :- http://IP:PORT/tasks
		body :- {
    			"title" : "job-1",
    			"description":"this is test job",
   				"due_date":"12-06-2035",
  				"status":"InProgress"
				}
	2. Get by the following APIs
		1. Get of all data :- http://IP:PORT/tasks/get
		2. Get using the title :- http://IP:PORT/tasks/get?title=job-103
		3. Get all tasks having status :- http://IP:PORT/tasks/get?status=Todo
	3. Update by following APIs 
		1. PUT APIs : - http://IP:PORT/tasks/update
	4. Delete by following APIs
		1. DELETE APIs :- http://IP:PORT/tasks/delete?title=job-1