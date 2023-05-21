CREATE table skills(
    sk_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    sk_name varchar,
    sk_description varchar,
    sk_group varchar,
    sk_category varchar,
    sk_purpose varchar,
    sk_priority int,
    sk_level int
);

INSERT INTO skills(sk_name, sk_description, sk_group, sk_category, sk_purpose, sk_priority, sk_level)
VALUES
('golang','Masterize go programming language', 'programming', 'language', 'general', 1, 3),
('python','Masterize python programming language', 'programming', 'language', 'general', 2, 7);
