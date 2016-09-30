<?php

$config = [
  'target' => [
    '/path/to/file1',
    '/path/to/file2',
    '/path/to/dir1', 
    '/path/to/dir2', 
  ],
  'log' => '/path/to/log',
  'port' => 3000,
];

echo json_encode($config), "\n";

