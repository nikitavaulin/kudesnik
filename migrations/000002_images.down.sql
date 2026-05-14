DROP INDEX IF EXISTS kudesnik.idx_products_has_image;
DROP INDEX IF EXISTS kudesnik.idx_products_missing_image;

ALTER TABLE kudesnik.products 
DROP COLUMN IF EXISTS thumbnail_url,
DROP COLUMN IF EXISTS image_url;