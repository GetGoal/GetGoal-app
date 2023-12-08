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
-- TOC entry 4841 (class 0 OID 16677)
-- Dependencies: 222
-- Data for Name: label; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.label (label_name) VALUES ('Exercise');
INSERT INTO public.label (label_name) VALUES ('Morning Activities');
INSERT INTO public.label (label_name) VALUES ('Self Development');
INSERT INTO public.label (label_name) VALUES ('Clean Food');


--
-- TOC entry 4835 (class 0 OID 16565)
-- Dependencies: 216
-- Data for Name: program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Morning Exercise Routine', 4.8, 'http://example.com/morning_exercise', 'A program for daily morning exercise routine', '30 minutes');
INSERT INTO public.program (program_name, rating, media_url, program_description, expected_time) VALUES ('Clean Eating Menu Plan', 5, 'http://example.com/clean_eating_menu', 'A program to plan a clean food menu for 7 days', '1 week');


--
-- TOC entry 4842 (class 0 OID 16682)
-- Dependencies: 223
-- Data for Name: label_program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.label_program (label_id, program_id) VALUES (1, 1);
INSERT INTO public.label_program (label_id, program_id) VALUES (2, 1);
INSERT INTO public.label_program (label_id, program_id) VALUES (4, 2);


--
-- TOC entry 4837 (class 0 OID 16663)
-- Dependencies: 218
-- Data for Name: user_account; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_account (first_name, last_name, email) VALUES ('Kellen', 'Heintze', 'kheintze0@gg.com');
INSERT INTO public.user_account (first_name, last_name, email) VALUES ('John', 'Doe', 'john.doe@gg.com');
INSERT INTO public.user_account (first_name, last_name, email) VALUES ('Val', 'Binden', 'vbinden3@gg.com');
INSERT INTO public.user_account (first_name, last_name, email) VALUES ('Dion', 'Donaho', 'ddonaho4@gg.com');


--
-- TOC entry 4839 (class 0 OID 16669)
-- Dependencies: 220
-- Data for Name: task; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Stretching', 1, 1, 0, '2023-11-15 06:00:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Perform stretching exercises', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Cardio Workout', 1, 1, 0, '2023-11-15 06:15:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Do cardio exercises like jogging or jumping jacks', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Strength Training', 1, 1, 0, '2023-11-15 06:30:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Include strength training exercises', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Cool Down', 1, 1, 0, '2023-11-15 06:45:00', NULL, 1, 'Morning Exercise Routine', NULL, 'Finish with cool-down and stretching', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 1: Breakfast', 1, 3, 0, '2023-10-25 08:00:00', NULL, 2, 'Clean Food', NULL, '1 cup low-fat plain Greek yogurt (179 calories)
1/4 cup raspberries (19 calories)
3 Tbsp. chopped walnuts (292 calories)', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 1: Lunch', 1, 3, 0, '2023-10-25 12:00:00', NULL, 2, 'Clean Food', NULL, '1 serving White Bean & Veggie Salad', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 1: Dinner', 1, 3, 0, '2023-10-25 18:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Sheet-Pan Roasted Salmon & Vegetables', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 2: Breakfast', 1, 3, 0, '2023-10-26 08:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Spinach, Peanut Butter & Banana Smoothie', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 2: Lunch', 1, 3, 0, '2023-10-26 12:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Vegan Superfood Grain Bowls', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 2: Dinner', 1, 3, 0, '2023-10-26 18:00:00', NULL, 2, 'Clean Food', NULL, '1 serving Slow-Cooker Vegetable Minestrone Soup (222 calories)
2 cups mixed greens (12 calories)
1/2 avocado, sliced (120 calories)
1 serving Citrus Vinaigrette (131 calories)', 'Null');
INSERT INTO public.task (task_name, task_status, user_account_id, is_set_noti, start_time, end_time, program_id, category, time_before_notify, task_description, link) VALUES ('Day 3: Breakfast', 1, 3, 0, '2023-10-27 08:00:00', NULL, 2, 'Clean Food', NULL, '1 cup low-fat plain Greek yogurt (179 calories)
1/4 cup raspberries (20 calories)
3 Tbsp. chopped walnuts (292 calories)', 'Null');
-- Continue with the remaining task INSERT statements


--
-- TOC entry 4846 (class 0 OID 16694)
-- Dependencies: 227
-- Data for Name: user_program; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (1, 1, 1);
INSERT INTO public.user_program (user_account_id, program_id, action_id) VALUES (3, 2, 1);

-- No need to set starting values for sequences in this script

COMMIT;
