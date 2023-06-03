# reflecton

Reflecton is a time keeping software that will help you to keep a track of your time in a much better way.

# Features 

# DB Architecture

## List of tables

Projects 
Tasks
Time Entries
Users
Roles

### Projects Tables 
Projects table will have the following columns
project_id
project_name
project_description
project_start_date
project_end_date
project_status
project_created_by
project_created_on
project_updated_by
project_updated_on
project_owner

### Users Table
Users table will have the following columns
user_id
user_name
user_email
user_password
user_created_by
user_created_on
user_updated_by
user_updated_on
user_role

# Roles table
Roles table will have the following columns
role_id
role_name
role_description
role_created_by
role_created_on
role_updated_by
role_updated_on


# Tasks Table
Tasks table will have the following columns
task_id
task_name 
task_type
task_description

# Time_entries Table
Time_entries table will have the following columns
time_entry_id
time_entry_date
time_entry_duration
is_prod_activity
is_billable
prod_servers
