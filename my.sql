--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg120+1)
-- Dumped by pg_dump version 15.3

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
-- Name: mydb; Type: DATABASE; Schema: -; Owner: leka
--

CREATE DATABASE mydb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE mydb OWNER TO leka;

\connect mydb

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
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: leka
--

INSERT INTO public.posts (id, author_id, title, content, created) VALUES (1, 2, 'USA kapzda', 'fuck you usa', '12:15:00');
INSERT INTO public.posts (id, author_id, title, content, created) VALUES (2, 3, 'dogs - super', 'take friend from pitompnik', '13:00:00');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: leka
--

INSERT INTO public.users (id, name) VALUES (1, 'Alexander');
INSERT INTO public.users (id, name) VALUES (2, 'Vasiliy');
INSERT INTO public.users (id, name) VALUES (3, 'Pyotr');
INSERT INTO public.users (id, name) VALUES (4, 'Simon');


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: leka
--

SELECT pg_catalog.setval('public.posts_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: leka
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- PostgreSQL database dump complete
--

