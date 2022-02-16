-- truncate current tables
-- TRUNCATE nft_txn_history;
-- TRUNCATE payment_history;
-- TRUNCATE chain_wallets;
-- TRUNCATE airdrop_campaigns;
-- TRUNCATE airdrop_campaigns, nft_collections;
-- TRUNCATE nft_txn_history, nfts;
-- TRUNCATE nft_txn_history, nfts, nft_collections;
-- TRUNCATE nft_txn_history, nfts, nft_collections, moments;
-- TRUNCATE nft_txn_history, nfts, nft_collections, moments, chain_wallets, payment_history, users;

-- insert several users
INSERT INTO users
(id, first_name, last_name, display_name, email_address, avatar_url, bio, youtube_url, tiktok_url, instagram_url)
VALUES
    (1, 'a', 'a', 'aa', 'a@gmail.com', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/uoDqE1bA.jpeg', 'I am aa', 'https://www.youtube.com/channel/UCsTeW2lMXnBSTM3spk1nu9w', 'https://www.tiktok.com/@kikakiim?lang=en', 'https://www.instagram.com/mayapolarbear/'),
    (2, 'b', 'b', 'bb', 'b@gmail.com', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/uoDqE1bA.jpeg', 'I am bb', 'https://www.youtube.com/watch?v=ipwxYa-F1uY', 'https://www.tiktok.com/@bts_official_bighit?is_copy_url=1&is_from_webapp=v1', 'https://www.instagram.com/samoyedpage/?hl=en');

ALTER SEQUENCE users_id_seq RESTART WITH 3;

-- insert chain wallets
INSERT INTO chain_wallets (id, owner_id, custodial_starkex_public_key, custodial_starkex_private_key, custodial_starkex_next_vault_id, custodial_ethereum_wallet_address, custodial_ethereum_private_key)
VALUES
    ('00000000-0000-0000-0000-000000002000', 1, '0x2deb04eb807be0ec943e08d8f666521edb3f12833922fbbc7f93e1434ae810e', '0x13c6ccabbd228c8f83cda0f20fb0ee639052c8afe75622c91b4a42dcf8da7ef', 3746326527, '0x0A0c2601C7874E77a401D91f8085DD07b040E595', '0xd5b63f3e2ddd3e94328f07e14be68dbb25b802bfa642643c517c349302a4f030'),
    ('00000000-0000-0000-0000-000000001000', 2, '0xf8c6635f9cfe85f46759dc2eebe71a45b765687e35dbe5e74e8bde347813ef', '0x607ba3969039f3e19006ff8f40629d20a7b7dac31d4019e0965fbf7c5c068a', 6738494583, '0x724f337bF0Fa934dB9aa3ec6dEa49B03c54AD3cc', '0xcfbac553dcb8e9c174bbeeb61fbf248e57504ec2c3e2e7a8dc9531f68f89536f');

-- insert moments
-- insert into moments (id, autograph_url, creator_id)
-- values
--     ('00000000-0000-0000-0000-000000000010', 'autograph_url1', 1),
--     ('00000000-0000-0000-0000-000000000020', 'autograph_url2', 2);

-- insert an nft_collection
insert into nft_collections
(id, collection_name, tags, nfts_amount, lowest_ask_price, highest_sale_price, description, loyalty_percentage, primary_sale_price, rarity)
values
    ('00000000-0000-0000-0000-000000000100', 'collection1', ARRAY['c1'],10, 1, 10, 'This is collection 1', 10, 40, 'Common'),
    ('00000000-0000-0000-0000-000000000200', 'collection2', ARRAY['c2'],20, 2, 20, 'This is collection 2', 20, 260, 'Uncommon');


-- insert several nfts
insert into nfts
(id, nft_collection_id, owner_id, serial_id, chain_type, txn_status, starkex_asset_id, ethereum_contract_address, ethereum_token_id, txn_id, txn_type, fixed_price, fixed_price_currency, onchain_txn_type, starkex_txn_id,media_type, media_url, cover_image_url,creator_id)
values
    ('00000000-0000-0000-0000-000000001000', '00000000-0000-0000-0000-000000000100', 1, 1, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0001-000000000000', 'FIXED_PRICE', 1, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),
    ('00000000-0000-0000-0000-000000002000', '00000000-0000-0000-0000-000000000100', 1, 2, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0002-000000000000', 'FIXED_PRICE', 2, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),
    ('00000000-0000-0000-0000-000000003000', '00000000-0000-0000-0000-000000000100', 1, 3, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0003-000000000000', 'FIXED_PRICE', 3, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),
    ('00000000-0000-0000-0000-000000004000', '00000000-0000-0000-0000-000000000100', 2, 4, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0004-000000000000', 'FIXED_PRICE', 4, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),
    ('00000000-0000-0000-0000-000000005000', '00000000-0000-0000-0000-000000000100', 2, 5, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0005-000000000000', 'FIXED_PRICE', 5, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),
    ('00000000-0000-0000-0000-000000006000', '00000000-0000-0000-0000-000000000100', 2, 6, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0006-000000000000', 'FIXED_PRICE', 6, 'USD', 'STARKEX', 'starkex_txn_id','VIDEO', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/gta5.mp4', 'https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fspecials-images.forbesimg.com%2Fimageserve%2F5ebe9c34f1f48800066306b0%2FGTA-Online%2F960x0.jpg%3Ffit%3Dscale',1),

    ('00000000-0000-0000-0000-000000007000', '00000000-0000-0000-0000-000000000200', 1, 1, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0007-000000000000', 'FIXED_PRICE', 2, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2),
    ('00000000-0000-0000-0000-000000008000', '00000000-0000-0000-0000-000000000200', 1, 2, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0008-000000000000', 'FIXED_PRICE', 2, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2),
    ('00000000-0000-0000-0000-000000009000', '00000000-0000-0000-0000-000000000200', 1, 3, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0009-000000000000', 'FIXED_PRICE', 3, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2),
    ('00000000-0000-0000-0000-000000010000', '00000000-0000-0000-0000-000000000200', 2, 4, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0010-000000000000', 'FIXED_PRICE', 4, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2),
    ('00000000-0000-0000-0000-000000011000', '00000000-0000-0000-0000-000000000200', 2, 5, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0011-000000000000', 'FIXED_PRICE', 5, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2),
    ('00000000-0000-0000-0000-000000012000', '00000000-0000-0000-0000-000000000200', 2, 6, 'NOT_ON_CHAIN', 'NOT_LISTED', 'starkex_asset_id', 'ethereum_contract_address', 'ethereum_token_id', '00000000-0000-0000-0012-000000000000', 'FIXED_PRICE', 6, 'USD', 'STARKEX', 'starkex_txn_id','IMAGE', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/ezgif-5-c6e4d685dc.mp4', 'https://storage.googleapis.com/creatornfts-bucket-dev/test-files/samoyed-dogs-puppies-1.jpeg',2);

insert into gifting_campaigns
(gifting_until_serial_id, nft_collection_id, creator_id)
values
    (10, '00000000-0000-0000-0000-000000000100', 1)