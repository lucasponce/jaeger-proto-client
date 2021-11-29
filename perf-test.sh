echo "Jaeger Perf Test"
for i in {1..50}
do
  ./status-test.sh &
done