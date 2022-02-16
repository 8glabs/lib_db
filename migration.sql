DROP TABLE IF EXISTS gifting_campaigns;
DROP TABLE IF EXISTS nft_txn_history;
DROP TABLE IF EXISTS payment_history;
DROP TABLE IF EXISTS chain_wallets;
DROP TABLE IF EXISTS nfts;
DROP TABLE IF EXISTS nft_collections;
-- DROP TABLE IF EXISTS moments;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS media_type_enum;
DROP TYPE IF EXISTS chain_type_enum;
DROP TYPE IF EXISTS txn_status_enum;
DROP TYPE IF EXISTS txn_type_enum;
DROP TYPE IF EXISTS onchain_txn_type_enum;
DROP TYPE IF EXISTS nft_lock_status_enum;
DROP TYPE IF EXISTS nft_collection_status_enum;
DROP TYPE IF EXISTS nft_collection_rarity_enum;

CREATE TYPE chain_type_enum AS ENUM('NOT_ON_CHAIN', 'ETHEREUM', 'STARKEX', 'SOLANA', 'LOOPRING');
CREATE TYPE txn_status_enum AS ENUM('NOT_LISTED', 'LISTED', 'PENDING', 'PAYMENT_MADE', 'FAILED');
CREATE TYPE txn_type_enum AS ENUM('FIXED_PRICE', 'FIRST_PRICE_AUCTION', 'SECOND_PRICE_AUCTION');
CREATE TYPE onchain_txn_type_enum AS ENUM('STARKEX', 'ETHEREUM');
CREATE TYPE nft_lock_status_enum AS ENUM('NOT_LOCKED', 'ADMIN_LOCKED', 'PRIMARY_SALE_LOCKED')
CREATE TYPE nft_collection_status_enum AS ENUM('GIFTING', 'PRIMARY_SALE', 'SECONDARY_SALE');
CREATE TYPE nft_collection_rarity_enum AS ENUM('Common', 'Uncommon', 'Rare', 'Epic', 'Mythic');

CREATE TABLE IF NOT EXISTS users (
                                     id BIGSERIAL NOT NULL PRIMARY KEY,
                                     first_name VARCHAR(100),
    last_name VARCHAR(100),
    display_name VARCHAR(250),
    hashed_password VARCHAR(255),
    phone_number VARCHAR(250),
    email_address VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255) DEFAULT '',
    bio VARCHAR(2000) DEFAULT '',
    youtube_url VARCHAR(255),
    tiktok_url VARCHAR(255),
    instagram_url VARCHAR(255),
    twitch_url VARCHAR(255),
    twitter_url VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email_address)
    );

-- A user may have many different walelts.
-- custodial_* means we create and keep the wallet for user, like Coinbase
-- In the future user may link their wallets, and we don't keep private keys
CREATE TABLE IF NOT EXISTS chain_wallets
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    owner_id BIGINT,
    custodial_starkex_public_key VARCHAR(255),
    custodial_starkex_private_key VARCHAR(255),
    custodial_starkex_next_vault_id BIGINT,
    custodial_ethereum_wallet_address VARCHAR(255),
    custodial_ethereum_public_key VARCHAR(255),
    custodial_ethereum_private_key VARCHAR(255),
    custodial_ethereum_mnemonics VARCHAR(255)[],
    CONSTRAINT fk_user_id FOREIGN KEY (owner_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION
    );


-- The media attached to an nft_collection is called moments
-- It contains a media, and may contain autograph, etc.
CREATE TYPE media_type_enum AS ENUM('VIDEO', 'IMAGE', 'GIF');

-- CREATE TABLE IF NOT EXISTS moments
-- (
--     id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
--     autograph_url VARCHAR(255),
--     CONSTRAINT fk_user_id FOREIGN KEY (creator_id)
--     REFERENCES users (id) MATCH SIMPLE
--                                 ON UPDATE CASCADE
--                                 ON DELETE NO ACTION
--     );

-- The nft collection created from the moments
CREATE TABLE IF NOT EXISTS nft_collections
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    collection_name VARCHAR(255) NOT NULL,
    tags VARCHAR(255)[],
    nfts_amount INTEGER DEFAULT 1,
    lowest_ask_price DOUBLE PRECISION,
    highest_sale_price DOUBLE PRECISION,
    description VARCHAR(2000),
    primary_sale_price INTEGER,
    -- constraint: between 0 and 100
    loyalty_percentage INTEGER,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status nft_collection_status_enum DEFAULT 'PRIMARY_SALE',
    rarity nft_collection_rarity_enum,
    
    );

-- The nfts in a collection
CREATE TABLE IF NOT EXISTS nfts
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    media_type media_type_enum NOT NULL,
    media_url VARCHAR(255) NOT NULL,
    cover_image_url VARCHAR(255),
    nft_collection_id uuid NOT NULL,
    owner_id BIGINT NOT NULL,
    buyer_id BIGINT,
    serial_id INTEGER NOT NULL,
    chain_type chain_type_enum NOT NULL,
    txn_status txn_status_enum NOT NULL,
    starkex_asset_id VARCHAR(255),
    starkex_vault_id BIGINT,
    ethereum_contract_address VARCHAR(255),
    ethereum_token_id VARCHAR(255),
    txn_id uuid UNIQUE,
    txn_type txn_type_enum,
    fixed_price DOUBLE PRECISION,
    fixed_price_currency VARCHAR(50),
    onchain_txn_type onchain_txn_type_enum,
    starkex_txn_id VARCHAR(255) DEFAULT '',
    stripe_payment_intent_id VARCHAR(255) DEFAULT '',
    nft_lock_status nft_lock_status_enum DEFAULT 'PRIMARY_SALE_LOCKED',
    media_upload_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- last_sale_price DOUBLE PRECISION,
    CONSTRAINT fk_user_id FOREIGN KEY (owner_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION,
    CONSTRAINT fk_collection_id FOREIGN KEY (nft_collection_id)
    REFERENCES nft_collections (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION
    CONSTRAINT fk_user_id FOREIGN KEY (creator_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION
    );

CREATE TABLE IF NOT EXISTS nft_txn_history
(
    id BIGSERIAL PRIMARY KEY,
    buyer_id BIGINT,
    seller_id BIGINT,
    nft_id uuid,
    txn_price DOUBLE PRECISION NOT NULL,
    txn_price_currency VARCHAR(255),
    txn_complete_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_nft_id FOREIGN KEY (nft_id)
    REFERENCES nfts (id) MATCH SIMPLE
                                ON UPDATE CASCADE
                                ON DELETE NO ACTION,
    CONSTRAINT fk_buyer_id FOREIGN KEY (buyer_id)
    REFERENCES users (id) MATCH SIMPLE
                                ON UPDATE CASCADE
                                ON DELETE NO ACTION,
    CONSTRAINT fk_seller_id FOREIGN KEY (seller_id)
    REFERENCES users (id) MATCH SIMPLE
                                ON UPDATE CASCADE
                                ON DELETE NO ACTION
    );

-- Keep a record of user payments
CREATE TABLE IF NOT EXISTS payment_history
(
    id BIGSERIAL PRIMARY KEY,
    payer_id BIGINT NOT NULL,
    stripe_order_id VARCHAR(255),
    amount DOUBLE PRECISION NOT NULL,
    currency VARCHAR(255),
    CONSTRAINT fk_buyer_id FOREIGN KEY (payer_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE NO ACTION
    );

-- Airdrop campaigns
CREATE TABLE IF NOT EXISTS gifting_campaigns
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    gifting_start_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    gifting_end_time TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    gifting_until_serial_id INTEGER NOT NULL,
    next_available_serial_id INTEGER NOT NULL DEFAULT 1,
    nft_collection_id uuid NOT NULL UNIQUE,
    creator_id BIGINT,
    CONSTRAINT fk_user_id FOREIGN KEY (creator_id)
    REFERENCES users (id) MATCH SIMPLE
                                 ON UPDATE CASCADE
                                 ON DELETE NO ACTION,
    CONSTRAINT fk_collection_id FOREIGN KEY (nft_collection_id)
    REFERENCES nft_collections (id) MATCH SIMPLE
                                 ON UPDATE CASCADE
                                 ON DELETE NO ACTION
    );