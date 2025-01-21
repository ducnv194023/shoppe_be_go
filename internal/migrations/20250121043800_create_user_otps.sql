-- +goose Up
CREATE TABLE user_otps (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL,
    user_id BIGINT NOT NULL,
    
    -- OTP info
    otp_code VARCHAR(6) NOT NULL,
    otp_type ENUM('login', 'reset_password', 'verify_email', 'verify_phone') NOT NULL,
    target VARCHAR(255) NOT NULL,  -- email hoặc phone
    
    -- Status & Security
    attempt_count INT DEFAULT 0,
    max_attempts INT DEFAULT 5,
    is_used BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP NOT NULL,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT fk_user_otps_user_id 
        FOREIGN KEY (user_id) 
        REFERENCES user_basics(id) 
        ON DELETE CASCADE,
    
    UNIQUE KEY unique_uuid (uuid),
    INDEX idx_user_id (user_id),
    INDEX idx_otp_code (otp_code),
    INDEX idx_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;