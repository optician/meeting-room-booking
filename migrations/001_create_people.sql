-- This is a sample migration.

create table meeting_rooms
(
	id uuid primary key, 
	name text unique, 
	capacity int, 
	office text, 
	stage int, 
	labels text[]
);
