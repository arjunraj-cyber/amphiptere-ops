Amphitheatre-Ops is a lightweight, unified Purple Team vulnerability scanner and defense verification utility that I wrote in Go. I wanted a simple tool that could bridge the gap between finding weaknesses from a Red Team perspective and instantly checking whether the Blue Team's logs, firewalls, and monitors are actually catching them all while generating a clean report at the end.
Features:
Runs simulated vulnerability scans and verifies corresponding defensive controls side-by-side.
Compiles everything neatly into a report.txt file so both teams have a record of what was found and what needs fixing.
Fast, clean, and has zero external dependencies to worry about.
