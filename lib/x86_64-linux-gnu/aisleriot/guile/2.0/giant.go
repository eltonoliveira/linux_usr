GOOF----LE-8-2.0À0      ] i 4     h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  ice-9¤	g  format¤	 ¤	 ¤	g  set-ace-low¤	g  
stock-slot¤								 ¤	g  
foundation¤				
						 ¤	g  tableau¤	g  reserve-slot¤	g  make-standard-double-deck¤	g  	make-deck¤	g  winning-score¤	g  allow-empty-slots¤	g  	same-suit¤	g  initialize-playing-area¤	g  shuffle-deck¤	g  add-normal-slot¤	 g  DECK¤	!g  stock¤	"g  add-blank-slot¤	#g  add-carriage-return-slot¤	$g  add-extended-slot¤	%g  down¤	&g  reserve¤	'g  deal-cards-face-up¤	(g  give-status-message¤	)e  4.5¤	*g  new-game¤	+g  set-statusbar-message¤	,g  get-stock-no-string¤	-g  _¤	.f  Deals left: ~a¤	/g  number->string¤	0g  length¤	1g  	get-cards¤	2g  member¤	3g  move-n-cards!¤	4g  reverse¤	5g  complete-transaction¤	6g  empty-slot?¤	7g  check-alternating-color-list¤	8g  check-same-suit-list¤	9g  check-straight-descending-list¤	:g  button-pressed¤	;g  	get-value¤	<g  ace¤	=g  get-suit¤	>g  get-top-card¤	?g  is-red?¤	@g  
droppable?¤	Ag  button-released¤	Bg  do-deal-next-cards¤	Cg  	dealable?¤	Dg  button-clicked¤	Eg  find-to¤	Fg  find-any-to-foundation¤	Gg  without-gaps¤	Hg  find-any-to-tableau¤	Ig  move-a-card¤	Jg  move-any-to-foundation¤	Kg  append¤	Lg  delayed-call¤	Mg  	auto-play¤	Ng  	add-card!¤	Og  remove-card¤	Pg  move-to-foundation¤	Qg  button-double-clicked¤	Rg  game-won¤	Sg  get-hint¤	Tg  	game-over¤	Ug  
game-score¤	Vg  
set-score!¤	Wg  any-slot-empty?¤	Xf  
Deal a row¤	Yg  	hint-move¤	Zf   Try moving a card to the reserve¤	[f  Try dealing a row of cards¤	\f  Try moving card piles around¤	]g  begin-exclusive¤	^f  	Same suit¤	_f  Alternating colors¤	`g  end-exclusive¤	ag  get-options¤	bg  list-ref¤	cg  apply-options¤	dg  timeout¤	eg  set-features¤	fg  droppable-feature¤	gg  dealable-feature¤	hg  
set-lambda¤C 5      h&    ] 4	 >  "  G  4i>   "  G  
RRR	R  h   Y   ] 6   Q       g  filenamef  	giant.scm
	
			 		
  g  nameg  	make-deck CR	hRRR !"#$%&'()       hø    ] 4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G  	
 C         g  filenamef  	giant.scm
	#
		$			%		#	&		3	)		9	)		>	)		G	*		W	+		Z	+		\	+		a	+		j	,		m	,		o	,		t	,		}	-	 	-	 	-	 	-	 	.	 	.	 	.	 	.	 £	/	 ¦	/	 ¨	/	 ­	/	 ¶	0	 ¹	0	 »	0	 À	0	 É	1	 Ì	1	 Î	1	 Ó	1	 Ü	2	 ß	2	 á	2	 æ	2	 ï	3	 ÿ	4		4		4		4		5		5		5	 	5	)	6	,	6	0	6	5	6	>	7	A	7	E	7	J	7	S	8	V	8	Z	8	_	8	h	9	k	9	o	9	t	9	}	:		:		:		:		;		;		;		;	§	<	·	=	º	=	¼	=	Á	=	Ê	?	Þ	B	ñ	C	ô	C	 S	õ
  g  nameg  new-game C*R+,   h   k   ] 45 6     c       g  filenamef  	giant.scm
	E
		F			F	 		
  g  nameg  give-status-message C(R-./01        h(   «   ] 4544455	56      £       g  filenamef  	giant.scm
	I
		K	
	
	K			K	
		L	
		L			L	%		L			L		 	L	
	"	J	 		"
  g  nameg  get-stock-no-string C,R234      h(   Ê   ]45$   456 6Â       g  
start-slot
		( g  	card-list		( g  end-slot			(  g  filenamef  	giant.scm
	R
		S			S			T	)		T		(	U	 		(	  g  nameg  complete-transaction C5R6789   hh   ö   ]
4 5$  "   $  C$  "  45$  "  $  45"  $  6C    î       g  slot
		d g  	card-list		d g  t				 g  t		6	X  g  filenamef  	giant.scm
	Y
		Z	
			Z			Z		"	Z		*	\		0	\	&	6	\		I	]		J	]	!	\	\		b	^	 		d	  g  nameg  button-pressed C:R2986;<=>74?0   h¸    ] $  C45$  r45$  b45$  R45$  45"  44544	55$  4544	55"  "  "  "  $  C4
5$  á45$  Ñ$  "  45$  "  $  45"  $  45$  "  445544	55$  `$  "  445544	55$  "  %$  445544	55"  "  "  "  "  $  C$  45$  
45CCC          g  
start-slot
	´ g  	card-list	´ g  end-slot		´ g  t	 ´ g  t	 Â ä g  t	=s g  t	´  g  filenamef  	giant.scm
	b
		c			c			d			d			e		%	d		&	f		0	d		1	g		;	g		<	h		A	h	$	C	h		F	h		K	i		P	i	(	R	i		S	i	9	V	i	C	^	i	9	_	i		c	i		d	j	!	i	j	,	k	j	!	l	j		m	j	A	p	j	L	x	j	A	y	j	 	d	 	n	 ¥	n	 ¦	o	 °	n	 ¶	p	 ¼	p	+ Â	p	 Õ	q	 Ö	q	& è	n	 é	r	 ó	r	 ù	s	! ü	s	1	s	,	s	!	s		s	K
	s	V	s	K	s		s		t	#	u	/&	u	?-	u	9/	u	/0	u	U3	u	^;	u	U<	u	)=	u	$=	t	Q	v	R	w	'U	w	6\	w	1^	w	'_	w	Lb	w	Vj	w	Lk	w	$	d		{		{		|	¦	{	§	}	¯	}	 R	´	  g  nameg  
droppable? C@R@5   h    ¹   ]4 5$  
 6C   ±       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  	giant.scm
 
	 		 		 	 			  g  nameg  button-released CAR'  h   d   ] 6      \       g  filenamef  	giant.scm
 
	
 	 		

  g  nameg  do-deal-next-cards CBRCB       h       ] $  45 $  6 CC           g  slot
		  g  filenamef  	giant.scm
 
	 		 		 	
	 		 	 		  g  nameg  button-clicked CDREF    h0   ö   ]	 &  C4 5$  	  C 6   î       g  
from-slots
		- g  find-to-result		-  g  filenamef  	giant.scm
 
	 			 		 		 	0	 		 		 		  		% 		+ 	$	- 	 		-  g  nameg  find-any-to-foundation CFR6G       h8   
  ] &  C$   C4 5$  	 6 4 5C      g  slots
		8 g  with-empties		8  g  filenamef  	giant.scm
 
	 			 		
 		 		 		 		 		  		% 	3	) 	%	, 		- 	#	2 	1	6 	#	7 	 		8	  g  nameg  without-gaps CGREGH019>6 
    h°     ] &  C445 5 "  	 6$  u4455$  "  44545 5$  64	5$  "  4455$   C"ÿÿ"ÿÿ|"ÿÿx       g  
from-slots
	 ¯ g  with-empties	 ¯ g  find-to-result		 ¯ g  cfs		 ¯ g  t		D	m g  t		y   g  filenamef  	giant.scm
 ¡
	 ¢			 ¢		 ¤		 ¤	%	 ¤	I	 ¤		 ¥		 ¤		, ®	!	0 ®		0 ¦		7 ¨		: ¨	 	B ¨		D ¨		D ¨		R ©		U ©	@	\ ©	Y	c ©	S	g ©	:	i ©		j ©		q ¦		r «		y «		y «	  ¬	  ¬	   ¬	  ¬	  ¦	 ¢ ­	 #	 ¯	  g  nameg  find-any-to-tableau CHRFI      h    Å   ]	4 5$  6C  ½       g  slots
		 g  find-any-result			  g  filenamef  	giant.scm
 °
	 ±			 ±		 ²		 ³		 ³	+	 ³	 		  g  nameg  move-any-to-foundation CJRJKLM        h       ] 44 55$  6C        g  filenamef  	giant.scm
 ·
	 ¸		 ¸	 	 ¸	0	 ¸	 	 ¸		 ¸		 ¹	 			
  g  nameg  	auto-play CMR6@>E        hH     ]
45$  "   $  C445  5$   C 6       g  slots
		H g  	from-slot		H g  t				  g  filenamef  	giant.scm
 ¿
	 À	
		 À		 À	-	 À	"	! À		$ Â		) Â	'	2 Â	!	5 Â	A	7 Â		; Â		> Ã		D Ä		H Ä	 		H	  g  nameg  find-to CER6NO   h(   ·   ]$  4 5"  $  C4 56 ¯       g  	from-slot
		' g  to-slot		'  g  filenamef  	giant.scm
 É
	 Ê			 Ê		 Ê		 Ê		 Ì		' Ì	 		'	  g  nameg  move-a-card CIRIE    h      ] 4 56       }       g  	from-slot
		  g  filenamef  	giant.scm
 Ð
	 Ñ		 Ñ	 		  g  nameg  move-to-foundation CPR2MP        h@   Ç   ]	4 5$  6 4 5$  "   $   6C   ¿       g  slot
		= g  t		1  g  filenamef  	giant.scm
 Õ
	 Ö		 Ö		 ×		 Ø		 Ø		. Ø	)	5 Ø		; Ù	 
		=  g  nameg  button-double-clicked CQR(RS    h(   v   ] 4>   "  G  45 $  C6        n       g  filenamef  	giant.scm
 à
	 á		 â		 â		! ã	 		!
  g  nameg  	game-over CTR01U     h(   Å   ] (  
C44 554 5C       ½       g  	slot-list
		!  g  filenamef  	giant.scm
 è
	 é		 ë			 ë		 ë		 ë		 ë			 ë	.	 ë	:	 ë	.	  ë	 		!  g  nameg  
game-score CURVU    h   u   ] 4455C      m       g  filenamef  	giant.scm
 ð
	 ñ		 ñ		 ñ		 ñ	 		
  g  nameg  game-won CRR6W-X      h@   ¼   ]45$  "    $  "  45 $  
45 CC  ´       g  t
		,  g  filenamef  	giant.scm
 ö
	 ø		 ÷		 ù		" ú		) ú		0 ÷		2 û		6 û		8 û		; û	 		>
  g  nameg  	dealable? CCRFKHY6-ZC[\     h°   6  ]44 5544 5544 55  $    6$  6$  645$  
4	5 C4
5 $  
45 C
45 C     .      g  find-result
	7 « g  	t-result1	7 « g  	t-result2		7 «  g  filenamef  	giant.scm

						-			=			-				
		
	+	 
	;	"
	+	%
		&		)	+	2	;	4	+	7		7			C		H		L	+	O		U		Z		^	)	a		g		l		p	)	s		t		~	 	0 	2 	0 	( 	 	 	! 	# 	! 	 ¡	 ¥	 §	 ª	 .	 «
  g  nameg  get-hint CSR]-^_`       h(   À   ] 45 45  C     ¸       g  filenamef  	giant.scm

																			&					"	 		#
  g  nameg  get-options CaRb     h      ]4 5 C       g  options
		  g  filenamef  	giant.scm

	 		 		 	 		  g  nameg  apply-options CcR    h   P   ] C    H       g  filenamef  	giant.scm
"
 		
  g  nameg  timeout CdR4eifigi>  "  G  hi*i:iAiDiQiTiRiSiaicidi@iCi6 ~      g  filenamef  	giant.scm		
		
	1	
	3			6	
	8			;	
	@	
 ´	
 ¹	
 ½	 
 Á	!
	#
	E
	I
¥	R

!	Y
	b
 
 
ê 
) 
 
j ¡
h °
< ·
« ¿
¢ É
X Ð
 Õ
9 à
@ è
è ð
! ö
$
%#
%Ò
&;"
&<$
&&
 )	&
   C6 