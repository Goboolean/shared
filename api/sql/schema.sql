-- metadata table for product data
CREATE TABLE product_meta (
  product_id    varchar(32) NOT NULL, /* an identifier form:{type}.{name}.{location}*/
  "name"        varchar(32) NOT NULL, /* human readable stock name */
  symbol        varchar(8)  NOT NULL, /* stock symbol */
  "description" text,                 /* product description */
  "type"        varchar(8)  NOT NULL, /* examples are stock, encrypt */
  exchange      varchar(32) NOT NULL, /* examples are kospi, nasdaq. */
  "location"    varchar(32),          /* examples are kor, usa. when product type is coin location is null*/
  PRIMARY KEY (product_id)
);

CREATE TABLE stock_platform (
  product_id varchar(32) NOT NULL, /* an product identifier form:{type}.{name}.{location}*/
  platform   varchar(16) NOT NULL, /* available platform is buycycle, polygon, kis */
  identifier varchar(32) NOT NULL, /* a string that is used to specific stock on such platform query */

  PRIMARY KEY (product_id),
  FOREIGN KEY (product_id) REFERENCES product_meta (product_id)
);

CREATE TABLE store_log (
  product_id  varchar(32) NOT NULL, /* an product identifier form:{type}.{name}.{location}*/
  stored_at   timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "status"    varchar(10) NOT NULL,

  PRIMARY KEY (stored_at),
  FOREIGN KEY (product_id) REFERENCES product_meta (product_id)
);
