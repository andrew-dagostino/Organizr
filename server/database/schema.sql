--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7
-- Dumped by pg_dump version 13.2

-- Started on 2021-05-19 14:33:45

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE organizr;
--
-- TOC entry 2918 (class 1262 OID 16634)
-- Name: organizr; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE organizr WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Canada.1252';


ALTER DATABASE organizr OWNER TO postgres;

\connect organizr

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16729)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 2919 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 204 (class 1259 OID 16742)
-- Name: board; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    title character varying(32) NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.board OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16740)
-- Name: board_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_id_seq OWNER TO postgres;

--
-- TOC entry 2920 (class 0 OID 0)
-- Dependencies: 203
-- Name: board_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_id_seq OWNED BY public.board.id;


--
-- TOC entry 210 (class 1259 OID 16785)
-- Name: board_member; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_member (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    member_id integer NOT NULL,
    board_id integer NOT NULL,
    board_permission_id integer NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.board_member OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16783)
-- Name: board_member_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_member_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_member_id_seq OWNER TO postgres;

--
-- TOC entry 2921 (class 0 OID 0)
-- Dependencies: 209
-- Name: board_member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_member_id_seq OWNED BY public.board_member.id;


--
-- TOC entry 206 (class 1259 OID 16755)
-- Name: board_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_permission (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(16) NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.board_permission OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16753)
-- Name: board_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_permission_id_seq OWNER TO postgres;

--
-- TOC entry 2922 (class 0 OID 0)
-- Dependencies: 205
-- Name: board_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_permission_id_seq OWNED BY public.board_permission.id;


--
-- TOC entry 208 (class 1259 OID 16768)
-- Name: member; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.member (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying(32) NOT NULL,
    email character varying(64) NOT NULL,
    password character varying(64) NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.member OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 16766)
-- Name: member_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.member_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.member_id_seq OWNER TO postgres;

--
-- TOC entry 2923 (class 0 OID 0)
-- Dependencies: 207
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.member_id_seq OWNED BY public.member.id;


--
-- TOC entry 214 (class 1259 OID 16833)
-- Name: task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    task_column_id integer NOT NULL,
    title character varying(32) NOT NULL,
    description character varying(256),
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.task OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16815)
-- Name: task_column; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_column (
    id integer NOT NULL,
    gid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    board_id integer NOT NULL,
    title character varying(32) NOT NULL,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.task_column OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16813)
-- Name: task_column_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_column_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_column_id_seq OWNER TO postgres;

--
-- TOC entry 2924 (class 0 OID 0)
-- Dependencies: 211
-- Name: task_column_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_column_id_seq OWNED BY public.task_column.id;


--
-- TOC entry 213 (class 1259 OID 16831)
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_seq OWNER TO postgres;

--
-- TOC entry 2925 (class 0 OID 0)
-- Dependencies: 213
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- TOC entry 2728 (class 2604 OID 16745)
-- Name: board id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board ALTER COLUMN id SET DEFAULT nextval('public.board_id_seq'::regclass);


--
-- TOC entry 2740 (class 2604 OID 16788)
-- Name: board_member id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member ALTER COLUMN id SET DEFAULT nextval('public.board_member_id_seq'::regclass);


--
-- TOC entry 2732 (class 2604 OID 16758)
-- Name: board_permission id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_permission ALTER COLUMN id SET DEFAULT nextval('public.board_permission_id_seq'::regclass);


--
-- TOC entry 2736 (class 2604 OID 16771)
-- Name: member id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.member ALTER COLUMN id SET DEFAULT nextval('public.member_id_seq'::regclass);


--
-- TOC entry 2748 (class 2604 OID 16836)
-- Name: task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- TOC entry 2744 (class 2604 OID 16818)
-- Name: task_column id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_column ALTER COLUMN id SET DEFAULT nextval('public.task_column_id_seq'::regclass);


--
-- TOC entry 2753 (class 2606 OID 16752)
-- Name: board board_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_gid_key UNIQUE (gid);


--
-- TOC entry 2769 (class 2606 OID 16795)
-- Name: board_member board_member_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_gid_key UNIQUE (gid);


--
-- TOC entry 2771 (class 2606 OID 16797)
-- Name: board_member board_member_member_id_board_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_member_id_board_id_key UNIQUE (member_id, board_id);


--
-- TOC entry 2773 (class 2606 OID 16793)
-- Name: board_member board_member_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_pkey PRIMARY KEY (id);


--
-- TOC entry 2757 (class 2606 OID 16765)
-- Name: board_permission board_permission_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_permission
    ADD CONSTRAINT board_permission_gid_key UNIQUE (gid);


--
-- TOC entry 2759 (class 2606 OID 16763)
-- Name: board_permission board_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_permission
    ADD CONSTRAINT board_permission_pkey PRIMARY KEY (id);


--
-- TOC entry 2755 (class 2606 OID 16750)
-- Name: board board_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_pkey PRIMARY KEY (id);


--
-- TOC entry 2761 (class 2606 OID 16782)
-- Name: member member_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT member_email_key UNIQUE (email);


--
-- TOC entry 2763 (class 2606 OID 16778)
-- Name: member member_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT member_gid_key UNIQUE (gid);


--
-- TOC entry 2765 (class 2606 OID 16776)
-- Name: member member_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT member_pkey PRIMARY KEY (id);


--
-- TOC entry 2767 (class 2606 OID 16780)
-- Name: member member_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.member
    ADD CONSTRAINT member_username_key UNIQUE (username);


--
-- TOC entry 2775 (class 2606 OID 16825)
-- Name: task_column task_column_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_column
    ADD CONSTRAINT task_column_gid_key UNIQUE (gid);


--
-- TOC entry 2777 (class 2606 OID 16823)
-- Name: task_column task_column_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_column
    ADD CONSTRAINT task_column_pkey PRIMARY KEY (id);


--
-- TOC entry 2779 (class 2606 OID 16843)
-- Name: task task_gid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_gid_key UNIQUE (gid);


--
-- TOC entry 2781 (class 2606 OID 16841)
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- TOC entry 2783 (class 2606 OID 16803)
-- Name: board_member board_member_board_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_board_id_fkey FOREIGN KEY (board_id) REFERENCES public.board(id) ON DELETE CASCADE;


--
-- TOC entry 2784 (class 2606 OID 16808)
-- Name: board_member board_member_board_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_board_permission_id_fkey FOREIGN KEY (board_permission_id) REFERENCES public.board_permission(id) ON DELETE CASCADE;


--
-- TOC entry 2782 (class 2606 OID 16798)
-- Name: board_member board_member_member_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_member
    ADD CONSTRAINT board_member_member_id_fkey FOREIGN KEY (member_id) REFERENCES public.member(id) ON DELETE CASCADE;


--
-- TOC entry 2785 (class 2606 OID 16826)
-- Name: task_column task_column_board_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_column
    ADD CONSTRAINT task_column_board_id_fkey FOREIGN KEY (board_id) REFERENCES public.board(id) ON DELETE CASCADE;


--
-- TOC entry 2786 (class 2606 OID 16844)
-- Name: task task_task_column_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_task_column_id_fkey FOREIGN KEY (task_column_id) REFERENCES public.task_column(id) ON DELETE CASCADE;


-- Completed on 2021-05-19 14:33:45

--
-- PostgreSQL database dump complete
--
