package auth

import "testing"

func TestPasswordRoundTrip(t *testing.T){encoded,err:=HashPassword("correct horse battery staple");if err!=nil{t.Fatal(err)};ok,err:=VerifyPassword(encoded,"correct horse battery staple");if err!=nil||!ok{t.Fatalf("expected password match: ok=%v err=%v",ok,err)};ok,err=VerifyPassword(encoded,"wrong");if err!=nil{t.Fatal(err)};if ok{t.Fatal("wrong password must not match")}}
func TestOpaqueTokensAreRandomAndHashable(t *testing.T){a,err:=NewOpaqueToken();if err!=nil{t.Fatal(err)};b,err:=NewOpaqueToken();if err!=nil{t.Fatal(err)};if a==b{t.Fatal("tokens must be unique")};if HashToken(a)==HashToken(b){t.Fatal("different tokens must have different hashes")}}
