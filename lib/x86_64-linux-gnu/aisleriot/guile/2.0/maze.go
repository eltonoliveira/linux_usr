GOOF----LE-8-2.0f      ] < 4  hά      ] g  guile€	 €	g  process-use-modules€	 €	 €	g  	aisleriot€	g  	interface€	 €		 €	
g  api€	
 €	 €	g  card€	g  add-normal-slot€	g  add-line€	g  initialize-playing-area€	g  set-ace-low€	g  make-standard-deck€	g  shuffle-deck€	g  DECK€	g  add-carriage-return-slot€	g  deal-cards-face-up€										
																					 	!	"	#	$	%	&	'	(	)	*	+	,	-	.	/	0	1	2	3	4	5 3€	g  make-visible-top-card€	g  eliminate-kings€	g  calculate-score€	g  new-game€	g  empty-slot?€	g  	get-value€	g  get-top-card€	g  king€	 g  remove-card€	!g  button-pressed€	"g  get-suit€	#g  
suit-next?€	$g  ace€	%g  queen€	&g  
card-next?€	'g  
droppable?€	(g  	add-card!€	)g  button-released€	*g  button-clicked€	+g  button-double-clicked€	,g  get-full-card€	-g  
set-score!€	.g  calculate-score-helper€	/g  	get-score€	0g  add-to-score!€	1g  game-won€	2g  	game-over€	3g  _€	4f  QAim to place the suits in the order which fits the current layout most naturally.€	5g  get-hint€	6g  get-options€	7g  apply-options€	8g  timeout€	9g  set-features€	:g  droppable-feature€	;g  
set-lambda€C 5       h(  ό   ] 4	 >  "  G  R     h   #  ] 4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  6       g  filenamef  maze.scm
	
																			%			(			-			6			9			>			G			J			O			X			[			`			i			l			q			z			}		 		 		 		 	 
  g  nameg  add-line CR    hΠ  ΰ  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4	

>  "  G  4
>  "  G  4	5>  "  G  4>   "  G  			 C      Ψ      g  filenamef  maze.scm
	 
		!			"		#	#		3	$		C	&		U	'		X	'		]	'		f	(		i	(		n	(		w	)		z	)			)	 	*	 	*	 	*	 	+	 	+	 ‘	+	 ͺ	,	 ­	,	 ²	,	 »	-	 Ύ	-	 Γ	-	 Μ	.	 Ο	.	 Τ	.	 έ	/	 ν	1	 ύ	2		3		4	-	5	=	6	M	7	]	8	m	9	}	;		;		;		@	‘	A	³	C	Ι	E	 0	Κ
  g  nameg  new-game CR      hP   Ι   ]4 5$  "  )44 55$  4 >  "  G  "    
$   6C     Α       g  slot
		K  g  filenamef  maze.scm
	G
		H			H			I	
		I			I	
	 	I		$	H		%	J		>	K		B	K		G	L		I	L	 		K  g  nameg  eliminate-kings CR  h      ]C    z       g  slot-id
		 g  	card-list		  g  filenamef  maze.scm
	N
 			  g  nameg  button-pressed C!R"   h(   Ή   ]4 545$  4 545CC±       g  first
		( g  second		(  g  filenamef  maze.scm
	Q
		R	
	
	S	
		R			R			T			T	
		U	
	%	T	 
		(	  g  nameg  
suit-next? C#R$%#        h8   ΐ   ]
45$  4 5"  $  C 6      Έ       g  lower
		2 g  higher		2 g  t			2  g  filenamef  maze.scm
	W
		X			X			X			Y			Y			X		2	Z	 			2	  g  nameg  
card-next? C&R'(     h    Ό   ]4 5$  	6C    ΄       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  maze.scm
	\
		]			]			^			^	 			  g  nameg  button-released C)R$&%       h   ―  ]  $  C45$  q
$  45"   45$  "  4455$  C	5$  45C45$  C456C§      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t		S   g  filenamef  maze.scm
	`
		a			a			b			b			c			b		!	d		%	d		(	e		/	e		4	f		9	f	&	;	f		?	f		E	g		H	g	 	M	g	.	O	g	 	S	g		S	d		c	h		g	h		j	i		q	i		s	j		x	j	&	z	j		~	j	 	k	% 	k	3 	k	% 	k	 "	 	  g  nameg  
droppable? C'R    h   k   ]C    c       g  slot-id
		  g  filenamef  maze.scm
	m
 		  g  nameg  button-clicked C*Rh   r   ]C    j       g  slot-id
		  g  filenamef  maze.scm
	p
 		  g  nameg  button-double-clicked C+R, h       ]4 5$   6 6             g  slot
		  g  filenamef  maze.scm
	s
		t			t			u			u			v	 		  g  nameg  get-full-card C,R,$-./    h`   ¦   ] 4
5 45$  4>  "  G  "  4
>  "  G  4>  "  G  45 	0C             g  filenamef  maze.scm
	x
		y		
	y			z				z			z			{		.	|		?	}		R	~		Y	~	 		Z
  g  nameg  calculate-score CR.&0       h   P  ]
 	6$  C4 5$  4 5"  $  C4 5 $  (45$  4>  "  G  "  "  $  C 6      H      g  slot
	  g  prev	  g  t		  g  t		,  g  t		n   g  filenamef  maze.scm
 
	 		 		 		 		 		# 	#	' 		, 		8 		@ 		E 		F 		R 		S 		n 		~ 	  	 	 	  g  nameg  calculate-score-helper C.R       h   Y   ] 6   Q       g  filenamef  maze.scm
 
	 	 		
  g  nameg  game-won C1R1      h   c   ] 45 C       [       g  filenamef  maze.scm
 
	 		 	 			
  g  nameg  	game-over C2R34  h   t   ] 
45 C  l       g  filenamef  maze.scm
 
	 	
	 		
 	
	 	 		
  g  nameg  get-hint C5R       h   S   ] C    K       g  filenamef  maze.scm
 
 		
  g  nameg  get-options C6Rh   k   ]C    c       g  options
		  g  filenamef  maze.scm
 
 		  g  nameg  apply-options C7Rh   O   ] C    G       g  filenamef  maze.scm
 
 		
  g  nameg  timeout C8R49i:i>  "  G  ;ii!i)i*i+i2i1i5i6i7i8i'i6  τ       g  filenamef  maze.scm		
					
ζ	
Γ	 
ό	G
	N
	Q
	£	W

	\
ϊ	`
~	m
	p
Κ	s
ρ	x
λ 
d 
ξ 
 
φ 
v 
Ϊ 
Ϋ 
& 
 	&
   C6 