grammar cmselect;

query : selectClause fromClause whereClause;

selectClause : SELECT attributeList;
attributeList : (attribute | '*') (',' (attribute | '*'))*;
attribute : IDENTIFIER;

fromClause : FROM tableName;
tableName : IDENTIFIER;

whereClause : WHERE condition (logicalOperator condition)*;
condition : attribute (EQUAL | NOT_EQUAL | LESS_THAN | GREATER_THAN | NOT_EQUAL_ALT) (NUMBER | STRING | parameter);
parameter : ':' IDENTIFIER ':';
logicalOperator : AND | OR | NOT;

EQUAL : '==';
NOT_EQUAL : '<>';
LESS_THAN : '<';
GREATER_THAN : '>';
NOT_EQUAL_ALT : '!=';

AND : 'AND';
OR : 'OR';
NOT : 'NOT';

SELECT : 'SELECT';
FROM : 'FROM';
WHERE : 'WHERE';

NUMBER : '-'? DIGIT+ ('.' DIGIT+)?;
STRING : '\'' (~'\'' | '\'\'' )* '\'';

IDENTIFIER : LETTER (LETTER | DIGIT | '_' | '.')*;

fragment LETTER : [a-zA-Z];
fragment DIGIT : [0-9];

WS : [ \t\r\n]+ -> skip;
