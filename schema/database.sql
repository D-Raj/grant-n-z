-- grant-n-z
DROP DATABASE IF EXISTS grant_n_z;
CREATE DATABASE IF NOT EXISTS grant_n_z;

-- use grant-n-z
USE grant_n_z;

-- If services exit, drop services
DROP TABLE IF EXISTS services;

-- If users exit, drop users
DROP TABLE IF EXISTS users;

-- If permissions exit, drop permissions
DROP TABLE IF EXISTS permissions;

-- If user_services exit, drop user_services
DROP TABLE IF EXISTS user_services;

-- If roles exit, drop roles
DROP TABLE IF EXISTS roles;

-- If policies exit, drop policies
DROP TABLE IF EXISTS policies;

-- services
CREATE TABLE services (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- users
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  username varchar(128) NOT NULL,
  email varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- permissions
CREATE TABLE permissions (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- roles
CREATE TABLE roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- role_members
CREATE TABLE role_members (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  user_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX (user_id),
  PRIMARY KEY (id),
  CONSTRAINT fk_role_members_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_role_members_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- user_services
CREATE TABLE user_services (
  id int(11) NOT NULL AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (user_id),
  INDEX (service_id),
  CONSTRAINT fk_user_services_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_user_services_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- policies
CREATE TABLE policies (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  permission_id int(11) NOT NULL,
  role_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (role_id),
  CONSTRAINT fk_policies_permission_id
  FOREIGN KEY (permission_id)
  REFERENCES permissions (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_policies_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;