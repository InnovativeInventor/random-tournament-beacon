cd std && gox
cd ..
cd simd && gox
cd ..

mv std/std* build/
mv simd/simd* build/
