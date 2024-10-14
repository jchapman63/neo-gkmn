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
-- Name: monster; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.monster (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    type text NOT NULL,
    basehp integer NOT NULL
);


--
-- Name: move; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.move (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    power integer NOT NULL,
    type text NOT NULL
);


--
-- Name: movemap; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movemap (
    monsterid uuid NOT NULL,
    moveid uuid NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: stats; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.stats (
    monsterid uuid NOT NULL,
    stattype text NOT NULL,
    power integer NOT NULL
);


--
-- Name: monster monster_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.monster
    ADD CONSTRAINT monster_pkey PRIMARY KEY (id);


--
-- Name: move move_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.move
    ADD CONSTRAINT move_pkey PRIMARY KEY (id);


--
-- Name: movemap movemap_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movemap
    ADD CONSTRAINT movemap_pkey PRIMARY KEY (monsterid, moveid);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: stats stats_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_pkey PRIMARY KEY (monsterid, stattype);


--
-- Name: movemap movemap_monsterid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movemap
    ADD CONSTRAINT movemap_monsterid_fkey FOREIGN KEY (monsterid) REFERENCES public.monster(id) ON DELETE CASCADE;


--
-- Name: movemap movemap_moveid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movemap
    ADD CONSTRAINT movemap_moveid_fkey FOREIGN KEY (moveid) REFERENCES public.move(id) ON DELETE CASCADE;


--
-- Name: stats stats_monsterid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_monsterid_fkey FOREIGN KEY (monsterid) REFERENCES public.monster(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20241014194139');
