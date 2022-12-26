--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5
-- Dumped by pg_dump version 12.5

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: players; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.players (
    name character varying NOT NULL,
    score smallint
);


ALTER TABLE public.players OWNER TO "user";

--
-- Data for Name: players; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.players (name, score) FROM stdin;
\.


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (name);


--
-- PostgreSQL database dump complete
--

