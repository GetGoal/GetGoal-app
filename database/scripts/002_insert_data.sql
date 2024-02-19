--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

-- Started on 2023-11-16 22:42:09
\c getgoal;

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4844 (class 0 OID 16688)
-- Dependencies: 225
-- Data for Name: action_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.action_type (action_name) VALUES ('Program owner');
INSERT INTO public.action_type (action_name) VALUES ('Join Program');
INSERT INTO public.action_type (action_name) VALUES ('Save Program');


--
-- TOC entry 4844 (class 0 OID 16688)
-- Dependencies: 225
-- Data for Name: email_valiadation_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.email_validation_status (status_description) VALUES ('Verified');
INSERT INTO public.email_validation_status (status_description) VALUES ('Pending');
INSERT INTO public.email_validation_status (status_description) VALUES ('Expired');
INSERT INTO public.email_validation_status (status_description) VALUES ('Invalid');


--
-- TOC entry 4841 (class 0 OID 16677)
-- Dependencies: 222
-- Data for Name: label; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.label (label_name) VALUES ('Exercise');
INSERT INTO public.label (label_name) VALUES ('Morning Activities');
INSERT INTO public.label (label_name) VALUES ('Self Development');
INSERT INTO public.label (label_name) VALUES ('Clean Food');
INSERT INTO public.label (label_name) VALUES ('Wellness');
INSERT INTO public.label (label_name) VALUES ('Mental well-being');
INSERT INTO public.label (label_name) VALUES ('Improve English');
INSERT INTO public.label (label_name) VALUES ('Vocabulary');
INSERT INTO public.label (label_name) VALUES ('Morning');
INSERT INTO public.label (label_name) VALUES ('Study');
INSERT INTO public.label (label_name) VALUES ('Tutorial');
INSERT INTO public.label (label_name) VALUES ('Beginners');
INSERT INTO public.label (label_name) VALUES ('Database');
INSERT INTO public.label (label_name) VALUES ('SQL');
INSERT INTO public.label (label_name) VALUES ('Relaxation');
INSERT INTO public.label (label_name) VALUES ('Sleep');


--
-- TOC entry 4835 (class 0 OID 16565)
-- Dependencies: 216
-- Data for Name: program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Morning Exercise Routine', 4.8, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F1-Morning-Exercise-Routine.jpg?alt=media&token=61d6ee0e-ccf2-4284-a985-41433569c472', 'A program for daily morning exercise routine', '30 minutes');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Clean Eating Menu Plan', 5, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F2-Clean-Eating-Menu-Plan.jpg?alt=media&token=56be62a0-c017-4c69-b39a-fd6c0b3471f5', 'A program to plan a clean food menu for 7 days', '3 days');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Full Day Wellness', 4.5, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F3-Full-Day-Wellness.jpg?alt=media&token=ad9975f8-8956-4115-8a15-16a17bc49dab', 'A holistic wellness program designed to promote physical and mental well-being throughout the day.', '1 day');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('How To Improve Your English Vocabulary in 7 Days', 4.4, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F4-How-To-Improve-Your-English-Vocabulary-in-7-Days.png?alt=media&token=9c6cd0a8-c74d-4bdb-8757-74757ce41188', 'In this lesson you will learn exactly how to improve your English vocabulary in 7 days. This method will not only help you learn more vocabulary, it will also help you to remember new vocabulary words and use them in sentences.', '1 week');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Morning rituals', 3.9, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F5-Morning-rituals.jpg?alt=media&token=fecd2e0d-927a-4fb5-ae02-e51ced17f509','', '1 day');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Study timer', 4, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F6-Study-timer.jpg?alt=media&token=39c543c3-e320-467a-b5f1-f0374edf784b' ,'Having a distinct study schedule helps you keep on track and not forget important aspects such as reviewing', '2 hours');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Learn Database Fundamentals', 3.2, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F7-Learn-Database-Fundamentals.jpg?alt=media&token=02ef4bc6-755e-4131-a3da-4289dc422462','Databases can be found in almost all software applications. SQL is the standard language to query a database. This SQL tutorial for beginners will teach you database design. Also, it teaches you basic to advanced SQL.', '3 days');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('SQL Basics', 5, 'https://firebasestorage.googleapis.com/v0/b/getgoal-2fccc.appspot.com/o/media%2Fprogram%2F8-SQL-Basics.jpg?alt=media&token=e257722f-4474-40ff-812e-6cc1d80ca170','Embark on a comprehensive learning journey to master the fundamentals of SQL (Structured Query Language), a powerful tool for managing and manipulating relational databases. ', '6 days');


--
-- TOC entry 4842 (class 0 OID 16682)
-- Dependencies: 223
-- Data for Name: label_program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.label_program (label_id, program_id) VALUES (1, 1);
INSERT INTO public.label_program (label_id, program_id) VALUES (2, 1);
INSERT INTO public.label_program (label_id, program_id) VALUES (4, 2);
INSERT INTO public.label_program (label_id, program_id) VALUES (5, 3);
INSERT INTO public.label_program (label_id, program_id) VALUES (1, 3);
INSERT INTO public.label_program (label_id, program_id) VALUES (6, 3);
INSERT INTO public.label_program (label_id, program_id) VALUES (7, 4);
INSERT INTO public.label_program (label_id, program_id) VALUES (8, 4);
INSERT INTO public.label_program (label_id, program_id) VALUES (2, 5);
INSERT INTO public.label_program (label_id, program_id) VALUES (9, 5);
INSERT INTO public.label_program (label_id, program_id) VALUES (10, 6);
INSERT INTO public.label_program (label_id, program_id) VALUES (11, 7);
INSERT INTO public.label_program (label_id, program_id) VALUES (12, 7);
INSERT INTO public.label_program (label_id, program_id) VALUES (13, 7);
INSERT INTO public.label_program (label_id, program_id) VALUES (14, 8);
INSERT INTO public.label_program (label_id, program_id) VALUES (13, 8);
INSERT INTO public.label_program (label_id, program_id) VALUES (11, 8);


--
-- TOC entry 4837 (class 0 OID 16663)
-- Dependencies: 218
-- Data for Name: user_account; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('Kellen', 'Heintze', 'kheintze0@gg.com','ssssss','bf',1);
INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('John', 'Doe', 'john.doe@gg.com','j123','bf',1);
INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('Val', 'Binden', 'vbinden3@gg.com','d123','bf',1);
INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('Dion', 'Donaho', 'ddonaho4@gg.com','t123','bf',1);
INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('Tiffani', 'Teach', 'tiffani@gg.com','tif123',  'bf',1);
INSERT INTO public.user_account (first_name, last_name, email ,password_hash ,password_salt ,email_validation_status_id) VALUES ('Test', '001', 'test001@gg.com','testja', 'bf',1);

--
-- TOC entry 4839 (class 0 OID 16669)
-- Dependencies: 220
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Stretching', 0, 1, 0, '2023-11-15 06:00:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Perform stretching exercises', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Cardio Workout', 0, 1, 0, '2023-11-15 06:15:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Do cardio exercises like jogging or jumping jacks', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Strength Training', 0, 1, 0, '2023-11-15 06:30:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Include strength training exercises', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Cool Down', 0, 1, 0, '2023-11-15 06:45:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Finish with cool-down and stretching', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 1: Breakfast', 0, 3, 0, '2023-10-25 08:00:00', NULL, 2, 'Clean Food', NULL, '1 cup low-fat plain Greek yogurt (179 calories)
1/4 cup raspberries (19 calories)
3 Tbsp. chopped walnuts (292 calories)', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 1: Lunch', 0, 3, 0, '2023-10-25 12:00:00', NULL, 2, 'Clean Food', NULL, '1 serving White Bean & Veggie Salad', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 1: Dinner', 0, 3, 0, '2023-10-25 18:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Sheet-Pan Roasted Salmon & Vegetables', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 2: Breakfast', 0, 3, 0, '2023-10-26 08:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Spinach, Peanut Butter & Banana Smoothie', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 2: Lunch', 0, 3, 0, '2023-10-26 12:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Vegan Superfood Grain Bowls', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 2: Dinner', 0, 3, 0, '2023-10-26 18:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Slow-Cooker Vegetable Minestrone Soup (222 calories)
2 cups mixed greens (12 calories)
1/2 avocado, sliced (120 calories)
1 serving Citrus Vinaigrette (131 calories)', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 3: Breakfast', 0, 3, 0, '2023-10-27 08:00:00', NULL, 2, 'Clean Food', NULL, '1 cup low-fat plain Greek yogurt (179 calories)
1/4 cup raspberries (20 calories)
3 Tbsp. chopped walnuts (292 calories)', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 3: Lunch', 0, 3, 0, '2023-10-27 12:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Vegan Superfood Grain Bowls', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 3: Dinner', 0, 3, 0, '2023-10-27 18:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Slow-Cooker Vegetable Minestrone Soup (222 calories)
2 cups mixed greens (12 calories)', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Morning Meditation', 0, 2, 0, '2023-11-02 07:00:00', NULL, 3, 'Full Day Wellness', NULL, 'Start your day with a meditation session to promote mindfulness and relaxation.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Healthy Breakfast', 0, 2, 0, '2023-11-02 08:00:00', NULL, 3, 'Full Day Wellness', NULL, 'Prepare and enjoy a nutritious breakfast to fuel your body for the day ahead.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Midday Stretch', 0, 2, 0, '2023-11-02 11:30:00', NULL, 3, 'Full Day Wellness', NULL, 'Take a break and do some stretching exercises to relieve tension and improve flexibility.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Healthy Lunch', 0, 2, 0, '2023-11-02 12:00:00', NULL, 3, 'Full Day Wellness', NULL, 'Prepare and enjoy a healthy lunch to nourish your body with essential nutrients.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Afternoon Walk', 0, 2, 0, '2023-11-02 14:00:00', NULL, 3, 'Full Day Wellness', NULL, 'Take a refreshing walk in the afternoon to boost energy and clear your mind.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Relaxation Session', 0, 2, 0, '2023-11-02 16:00:00', NULL, 3, 'Full Day Wellness', NULL, 'Wind down with a relaxation session to promote calmness and prepare for a restful night.', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 1 Find your interest', 0, 5, 0, '2023-11-05 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'https://www.youtube.com/watch?v=bl3UBU_nIPQ', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 2 Read the article or blog', 0, 5, 0, '2023-11-06 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 3 Write the definations and sentences', 0, 5, 0, '2023-11-07 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 4 Continue rite the definitions and sentences', 0, 5, 0, '2023-11-08 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 5 Find one image to match each word', 0, 5, 0, '2023-11-09 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 6 Make sentences about the images', 0, 5, 0, '2023-11-10 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Day 7 Review the images', 0, 5, 0, '2023-11-11 16:00:00', NULL, 4, 'How To Improve Your English Vocabulary in 7 Days', NULL, 'Null', 'Null', 'www.mockexample.com');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Make Bed', 0, 4, 0, '2023-11-15 07:00:00', NULL, 5, 'Morning rituals', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Meditation', 0, 4, 0, '2023-11-15 08:00:00', NULL, 5, 'Morning rituals', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Exercise', 0, 4, 0, '2023-11-15 09:00:00', NULL, 5, 'Morning rituals', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Drink Tea', 0, 4, 0, '2023-11-15 10:00:00', NULL, 5, 'Morning rituals', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Journal', 0, 4, 0, '2023-11-15 11:00:00', NULL, 5, 'Morning rituals', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Check the subject of study', 0, 2, 0, '2023-11-15 19:00:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Concentration', 0, 2, 0, '2023-11-15 19:30:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Rest', 0, 2, 0, '2023-11-15 20:30:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Concentration', 0, 2, 0, '2023-11-15 20:40:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Rest', 0, 2, 0, '2023-11-15 21:40:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('Revise', 0, 2, 0, '2023-11-15 22:00:00', NULL, 6, 'Study timer', NULL, 'Null', 'Null', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('What is a Database?', 0, 1, 0, '2023-11-20 20:00:00', NULL, 7, 'Learn Database Fundamentals', NULL, 'Definition, Meaning, Types with Example', 'https://www.guru99.com/introduction-to-database-sql.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('What is SQL?', 0, 1, 0, '2023-11-21 20:00:00', NULL, 7, 'Learn Database Fundamentals', NULL, 'Learn SQL Basics, SQL Full Form & How to Use', 'https://www.guru99.com/what-is-sql.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL Workbench Tutorial', 0, 1, 0, '2023-11-22 20:00:00', NULL, 7, 'Learn Database Fundamentals', NULL, 'What is, How to Install & Use', 'https://www.guru99.com/introduction-to-mysql-workbench.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL Create Table', 0, 1, 0, '2023-11-23 20:00:00', NULL, 8, 'SQL Basics', NULL, 'How to Create Database in MySQL', 'https://www.guru99.com/how-to-create-a-database.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL SELECT Statement', 0, 1, 0, '2023-11-24 20:00:00', NULL, 8, 'SQL Basics', NULL, 'Learn with Example', 'https://www.guru99.com/select-statement.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL WHERE Clause', 0, 1, 0, '2023-11-25 20:00:00', NULL, 8, 'SQL Basics', NULL, 'AND, OR, IN, NOT IN Query Example', 'https://www.guru99.com/where-clause.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL INSERT INTO Query', 0, 1, 0, '2023-11-26 20:00:00', NULL, 8, 'SQL Basics', NULL, 'How to add Row in Table (Example)', 'https://www.guru99.com/insert-into.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL DELETE Query', 0, 1, 0, '2023-11-27 20:00:00', NULL, 8, 'SQL Basics', NULL, 'How to Delete Row from a Table', 'https://www.guru99.com/insert-into.html', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link, media_url) VALUES ('MySQL UPDATE Query', 0, 1, 0, '2023-11-28 20:00:00', NULL, 8, 'SQL Basics', NULL, 'Learn with Example', 'https://www.guru99.com/sql-update-query.html', 'Null');
-- Continue with the remaining task INSERT statements


--
-- TOC entry 4846 (class 0 OID 16694)
-- Dependencies: 227
-- Data for Name: user_program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (1, 1, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (3, 2, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (2, 3, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (5, 4, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (4, 5, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (2, 6, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (1, 7, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (1, 8, 1);


-- No need to set starting values for sequences in this script

COMMIT;
