CREATE TABLE IF NOT EXISTS nft(
                                  owner_address VARCHAR(128) PRIMARY KEY,
                                  token_id VARCHAR(128) NOT NULL,
                                  name VARCHAR(128),
                                  description VARCHAR(2048),
                                  image_url VARCHAR(256),
                                  image_preview_url VARCHAR(256),
                                  image_thumbnail_url VARCHAR(256),
                                  contract_address VARCHAR(128),
                                  contract_name VARCHAR(128),
                                  contract_symbol VARCHAR(128),
                                  contract_description VARCHAR(2048),
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);