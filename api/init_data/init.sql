INSERT INTO wordtree_dev.users (name, password) VALUES
('user1', 'user1pass'),
('user2', 'user2pass'),
('user3', 'user3pass');

INSERT INTO wordtree_dev.word_definitions (word_name, meaning, keyword) VALUES
('apple', 'a round fruit with firm, white flesh and a green, red, or yellow skin', ''),
('flesh', 'the soft substance consisting of muscle and fat that is found between the skin and bones of a human or an animal', ''),
('muscle', 'physical power; strength.', ''),
('substance', 'a particular kind of matter with uniform properties.', '');

INSERT INTO wordtree_dev.word_nodes (uuid, user_id, word_definition_id) VALUES
('user1', 1, 1),
('user1', 1, 2),
('user1', 1, 3),
('user1', 1, 4);

INSERT INTO wordtree_dev.current_node_relations (current_node_id, left_node_id, right_node_id) VALUES
(1, null, null),
(1, null, null),
(1, null, null);

INSERT INTO wordtree_dev.word_node_childrens (current_node_id, child_node_id) VALUES
(1, 2),
(2, 3),
(2, 4);