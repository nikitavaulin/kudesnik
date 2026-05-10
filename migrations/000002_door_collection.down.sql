UPDATE kudesnik.doors SET collection = '' WHERE collection IS NULL;
ALTER TABLE kudesnik.doors ALTER COLUMN collection SET NOT NULL;