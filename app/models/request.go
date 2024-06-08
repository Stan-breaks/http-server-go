package models

type Request struct{
  Method  string
  Path  string
  Headers  map[string]string
  Body  string
}
