#!/bin/bash	
 bcHeight=0	
function setBcHeight()	
{	
	local -n ref=$1	
	bcInfo=$(./cli getBlockchainInfo)	
	arr=(${bcInfo// / })	
	height=${arr[-1]}	
	if [[ $bcInfo == *"ERROR: GetBlockchainInfo failed."* ]]; then	
  		echo "Error has occurred, make sure dapp service is running."; exit 1	
	fi	
	ref=$height	
}	
 # set current blockchain height to bcHeight	
setBcHeight bcHeight	
 # declare some variables	
newBcHeight=0	
addBalanceAmount=10	
# check input argument exists	
if [ -z "$1" ]	
  then	
    echo "Error: Missing argument. How many wallets do you want?"; exit 1	
fi	
# check input argument $1 is number	
re='^[0-9]+$'	
if ! [[ $1 =~ $re ]] ; then	
   echo "Error: Not a number"; exit 1	
fi	
# check input arg $1 >= 2	
if [ $1 -le 1 ];  then	
	echo "Error: At least 2 wallets should be created"; exit 1	
fi	
 # check input arg $2 >= 1	
if [ $2 -le 1 ];  then	
	echo "Error: Too few transactions"; exit 1	
fi	
# loop $1 times	
for (( c=0; c<=$1; c++ )); do	
	# create wallet, respond to every command line prompt with 'y' and store print information	
	output=$(yes | ./cli createWallet)	
	while read -r line; do	
		arr=(${line// / })	
		# get address in output	
		address=${arr[-1]}	
		# if is not address (operation failed) then create wallet again	
		if [ ${#address} -ne "34" ]; then	
			./createWallets $1 $2	
			exit 1	
		fi	
		# create list of addresses	
		accList[$c]=$address		
	done <<< "$output"	
done	
 function reviewBalancesAndQuit()	
{	
	for i in "${accList[@]}"	
	do 	
	line=$(./cli getBalance -address $i)	
	arr=(${line// / })		
	amount=${arr[-1]}	
	echo "$i balance: $amount"	
	done	
	echo "Job done."	
	exit 1	
}	
counter=0	
while [ $counter -le $2 ]; do	
	if [ $bcHeight -lt $newBcHeight ]; then	
		if [ $counter -eq $2 ]; then	
			reviewBalancesAndQuit	
		fi	
		echo "block mined...height $newBcHeight"	
		# miner gives some money	
	        address=${accList[$RANDOM % ${#accList[@]} ]}		
		echo "adding $addBalanceAmount to $address"	
		./cli addBalance -address $address -amount $addBalanceAmount	
		#if [[ $addBalanceRes == *"Add balance error!"* ]]; then	
		#	echo $addBalanceRes; exit 1	
   		#fi	
		# emulate a random transaction	
		amount=$((RANDOM%(addBalanceAmount+10)))	
		from=$address	
		to=${accList[$RANDOM % ${#accList[@]} ]} 	
		echo "sending $amount from $from to $to"	
		./cli send -from $from -to $to -amount $amount # &>> log	
		bcHeight=$newBcHeight	
       		((counter++))		
	       echo "waiting for miner to mint next block..."	
	       echo "---------------------------"	
	else	
	       setBcHeight newBcHeight	
	       sleep 1	
       fi	
done	
