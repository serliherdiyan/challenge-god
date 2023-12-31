PGDMP     :        	            {            enigma_laundry    15.3    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16755    enigma_laundry    DATABASE     �   CREATE DATABASE enigma_laundry WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE enigma_laundry;
                postgres    false            �            1259    16757    customer    TABLE     �   CREATE TABLE public.customer (
    customer_id integer NOT NULL,
    customer_name character varying(50) NOT NULL,
    phone_number character varying(15) NOT NULL
);
    DROP TABLE public.customer;
       public         heap    postgres    false            �            1259    16756    customer_customer_id_seq    SEQUENCE     �   CREATE SEQUENCE public.customer_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.customer_customer_id_seq;
       public          postgres    false    215                       0    0    customer_customer_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.customer_customer_id_seq OWNED BY public.customer.customer_id;
          public          postgres    false    214            �            1259    16764    service    TABLE     �   CREATE TABLE public.service (
    service_id integer NOT NULL,
    service_name character varying(50) NOT NULL,
    service_unit character varying(10) NOT NULL,
    price_per_unit integer NOT NULL
);
    DROP TABLE public.service;
       public         heap    postgres    false            �            1259    16763    service_service_id_seq    SEQUENCE     �   CREATE SEQUENCE public.service_service_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.service_service_id_seq;
       public          postgres    false    217                       0    0    service_service_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.service_service_id_seq OWNED BY public.service.service_id;
          public          postgres    false    216            �            1259    16771    transaction    TABLE     >  CREATE TABLE public.transaction (
    transaction_id integer NOT NULL,
    customer_id integer NOT NULL,
    service_id integer NOT NULL,
    quantity integer NOT NULL,
    total_price integer NOT NULL,
    entry_date date NOT NULL,
    completion_date date NOT NULL,
    received_by character varying(50) NOT NULL
);
    DROP TABLE public.transaction;
       public         heap    postgres    false            �            1259    16770    transaction_transaction_id_seq    SEQUENCE     �   CREATE SEQUENCE public.transaction_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 5   DROP SEQUENCE public.transaction_transaction_id_seq;
       public          postgres    false    219                       0    0    transaction_transaction_id_seq    SEQUENCE OWNED BY     a   ALTER SEQUENCE public.transaction_transaction_id_seq OWNED BY public.transaction.transaction_id;
          public          postgres    false    218            o           2604    16760    customer customer_id    DEFAULT     |   ALTER TABLE ONLY public.customer ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_customer_id_seq'::regclass);
 C   ALTER TABLE public.customer ALTER COLUMN customer_id DROP DEFAULT;
       public          postgres    false    215    214    215            p           2604    16767    service service_id    DEFAULT     x   ALTER TABLE ONLY public.service ALTER COLUMN service_id SET DEFAULT nextval('public.service_service_id_seq'::regclass);
 A   ALTER TABLE public.service ALTER COLUMN service_id DROP DEFAULT;
       public          postgres    false    217    216    217            q           2604    16774    transaction transaction_id    DEFAULT     �   ALTER TABLE ONLY public.transaction ALTER COLUMN transaction_id SET DEFAULT nextval('public.transaction_transaction_id_seq'::regclass);
 I   ALTER TABLE public.transaction ALTER COLUMN transaction_id DROP DEFAULT;
       public          postgres    false    218    219    219            	          0    16757    customer 
   TABLE DATA           L   COPY public.customer (customer_id, customer_name, phone_number) FROM stdin;
    public          postgres    false    215   �                 0    16764    service 
   TABLE DATA           Y   COPY public.service (service_id, service_name, service_unit, price_per_unit) FROM stdin;
    public          postgres    false    217   !                  0    16771    transaction 
   TABLE DATA           �   COPY public.transaction (transaction_id, customer_id, service_id, quantity, total_price, entry_date, completion_date, received_by) FROM stdin;
    public          postgres    false    219                      0    0    customer_customer_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.customer_customer_id_seq', 7, true);
          public          postgres    false    214                       0    0    service_service_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.service_service_id_seq', 11, true);
          public          postgres    false    216                       0    0    transaction_transaction_id_seq    SEQUENCE SET     L   SELECT pg_catalog.setval('public.transaction_transaction_id_seq', 1, true);
          public          postgres    false    218            s           2606    16762    customer customer_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);
 @   ALTER TABLE ONLY public.customer DROP CONSTRAINT customer_pkey;
       public            postgres    false    215            u           2606    16769    service service_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_pkey PRIMARY KEY (service_id);
 >   ALTER TABLE ONLY public.service DROP CONSTRAINT service_pkey;
       public            postgres    false    217            w           2606    16776    transaction transaction_pkey 
   CONSTRAINT     f   ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (transaction_id);
 F   ALTER TABLE ONLY public.transaction DROP CONSTRAINT transaction_pkey;
       public            postgres    false    219            x           2606    16777 (   transaction transaction_customer_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id);
 R   ALTER TABLE ONLY public.transaction DROP CONSTRAINT transaction_customer_id_fkey;
       public          postgres    false    215    3187    219            y           2606    16782 '   transaction transaction_service_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.service(service_id);
 Q   ALTER TABLE ONLY public.transaction DROP CONSTRAINT transaction_service_id_fkey;
       public          postgres    false    217    219    3189            	   >   x�3�H̫�4�053 �2��H�+����%ps���d�$�EM� �6F��� ;8�         N   x�3�tu���v�4600�2�v	��v���D��
\����AR&�>��y)E�
N�)��e�E )S�\� ��<         9   x�3�4�4�4�44 N##c]s]#c8�؀�8�('���ԘӔ�Д��=... ֻ�     