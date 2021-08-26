rm ./t1c
rm ./t1u
./sort text > t1c

sort text > t1u
diff t1c t1u