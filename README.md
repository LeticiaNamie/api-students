# api-students
API to manage students, made in 'Golang do Zero' course

Routes:
-GET /students - List all students
-POST /students - Create new student
-GET /students/:id - Get infos from a specific student
-PUT /students/:id - Update student
-DELETE /students/:id - Delete student
-GET /students?active=<true/false> - List of all active/non-active students

Struct Student
-Name (string)
-CPF (int)
-Email (string)
-Age (int)
-Active (bool)
