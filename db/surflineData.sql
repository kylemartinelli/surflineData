DROP DATABASE IF EXISTS surflineData;
CREATE DATABASE surflineData;

USE surflineData;

DROP TABLE IF EXISTS spots;

CREATE TABLE IF NOT EXISTS spots(
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `surflineId` INT UNSIGNED NOT NULL,
  `surflineUrl` VARCHAR(255) NOT NULL,
  `spotName` VARCHAR(255) NOT NULL,
  `lat` DECIMAL(8,6),
  `lon` DECIMAL(9,6),
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  UNIQUE KEY (surflineId)
);

DROP TABLE IF EXISTS idealConditions;

CREATE TABLE IF NOT EXISTS idealConditions(
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `spotId` INT UNSIGNED NOT NULL,
  `idealSeason` VARCHAR(255) NOT NULL,
  `idealSeasonDescription` VARCHAR(255) NOT NULL,
  `idealTide` VARCHAR(255) NOT NULL,
  `idealTideDescription` VARCHAR(255) NOT NULL,
  `idealSize` VARCHAR(255) NOT NULL,
  `idealSizeDescription` VARCHAR(255) NOT NULL,
  `idealWindDirections` VARCHAR(255) NOT NULL,
  `idealWindDirectionsDescription` VARCHAR(255) NOT NULL,
  `idealSwellDirections` VARCHAR(255) NOT NULL,
  `idealSwellDirectionsDescription` VARCHAR(255) NOT NULL,
  `abilityLevels` VARCHAR(255) NOT NULL,
  `abilityDescription` VARCHAR(255) NOT NULL,
  `bottomType` VARCHAR(255) NOT NULL,
  `bottomTypeDescription` VARCHAR(255) NOT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(spotId) REFERENCES spots(id) ON DELETE CASCADE
);


DROP TABLE IF EXISTS spotCharacteristics;
CREATE TABLE IF NOT EXISTS spotCharacteristics(
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `spotId` INT UNSIGNED NOT NULL,
  `crowdFactorDescription` VARCHAR(255) NOT NULL,
  `crowdFactorRating` INT UNSIGNED NOT NULL,
  `crowdFactorSummary` VARCHAR(255) NOT NULL,

  `localVibeDescription` VARCHAR(255) NOT NULL,
  `localVibeRating` INT UNSIGNED NOT NULL,
  `localVibeSummary` VARCHAR(255) NOT NULL,

  `shoulderBurnDescription` VARCHAR(255) NOT NULL,
  `shoulderBurnRating` INT UNSIGNED NOT NULL,
  `shoulderBurnSummary` VARCHAR(255) NOT NULL,

  `spotRatingDescription` VARCHAR(255) NOT NULL,
  `spotRatingRating` INT UNSIGNED NOT NULL,
  `spotRatingSummary` VARCHAR(255) NOT NULL,

  `waterQualityDescription` VARCHAR(255) NOT NULL,
  `waterQualityRating` INT UNSIGNED NOT NULL,
  `waterQualitySummary` VARCHAR(255) NOT NULL,

  `boardTypes`  VARCHAR(255) NOT NULL,

  `breakType`  VARCHAR(255) NOT NULL,
  `spotDescription` TEXT,


  PRIMARY KEY(id),
  FOREIGN KEY (spotId) REFERENCES spots(id)

);


DROP TABLE IF EXISTS spotBreadCrumbs;

CREATE TABLE IF NOT EXISTS spotBreadCrumbs(
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `spotId` INT UNSIGNED NOT NULL,
  `country` VARCHAR(255) NOT NULL,
  `countryId` INT UNSIGNED NOT NULL,
  `region` VARCHAR(255) NOT NULL,
  `regionId` INT UNSIGNED NOT NULL,
  `subRegion` VARCHAR(255) NOT NULL,
  `subRegionId` INT UNSIGNED NOT NULL,
  `area` VARCHAR(255) DEFAULT NULL,
  `areaId` INT UNSIGNED DEFAULT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY(spotId) REFERENCES spots(id)
);


-- mysql -u root -h localhost -p < db/surflineData.sql