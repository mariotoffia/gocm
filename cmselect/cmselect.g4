grammar cmselect;

query : selectClause fromClause whereClause;

selectClause : SELECT attributeList;
attributeList : attribute (',' attribute)*;
attribute : (IDENTIFIER | '*') (AS IDENTIFIER)?;
rhAttribute: IDENTIFIER;

fromClause : FROM tableName;
tableName : IDENTIFIER;

whereClause : WHERE condition (logicalOperator condition)*;
condition : attribute (EQUAL | NOT_EQUAL | LESS_THAN | GREATER_THAN) rightHandExpression;
rightHandExpression :  STRING | NUMBER | parameter| rhAttribute;

parameter : ':' IDENTIFIER ':';

logicalOperator : AND | OR | NOT;

EQUAL : '==';
NOT_EQUAL : '!=';
LESS_THAN : '<';
GREATER_THAN : '>';

AND : 'AND';
OR : 'OR';
NOT : 'NOT';

AS : 'AS';

SELECT : 'SELECT';
FROM : 'FROM';
WHERE : 'WHERE';

NUMBER : '-'? DIGIT+ ('.' DIGIT+)?;

// Add support for double-quoted strings
STRING : ('\'' (~'\'' | '\'\'' )* '\'') | ('"' (~'"' | '""' )* '"');

IDENTIFIER : LETTER (LETTER | DIGIT | '_' | '.')*;

fragment LETTER : [a-zA-Z];
fragment DIGIT : [0-9];

WS : [ \t\r\n]+ -> skip;
