ALTER TABLE kudesnik.products 
ADD COLUMN IF NOT EXISTS image_url TEXT,
ADD COLUMN IF NOT EXISTS thumbnail_url TEXT;


CREATE INDEX IF NOT EXISTS idx_products_has_image ON kudesnik.products (image_url) 
WHERE image_url IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_products_missing_image ON kudesnik.products (product_id) 
WHERE image_url IS NULL;