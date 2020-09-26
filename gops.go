// Gops implements this at the moment with code in goprocess.FindAll() that looks like this, in somewhat sketched and reduced form:
func FindAll() []P {
   pss, err := ps.Processes()
   [...]
   found := make(chan P)
   limitCh := make(chan struct{}, concurrencyProcesses)

   for _, pr := range pss {
      limitCh <- struct{}{}
      pr := pr
      go func() {
         defer func() { <-limitCh }()
         [... get a P with some error checking ...]
         found <- P
      }()
   }
   [...]

   var results []P
   for p := range found {
      results = append(results, p)
   }
   return results
}
/*
The bug is that the goroutines only receive from limitCh to release their token after sending their result to the
unbuffered found channel, while the main code only starts receiving from found after running through the entire
loop, and the main code takes the token in the loop and blocks if no tokens are available. So if you have too many
processes to go through, you start N goroutines, they all block trying to write to found and don't receive from
limitCh, and the main for loop blocks trying to send to limitCh and never reaches the point where it starts
receiving from found.
*/
