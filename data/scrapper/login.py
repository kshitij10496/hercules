import requests
from bs4 import BeautifulSoup
import getpass
import time


ERP_LOGIN_URL = 'https://erp.iitkgp.ac.in/IIT_ERP3/' #to get the session cookie
SECURITY_QUESTION_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/getSecurityQues.htm'


while (1):
    roll_no = input("Enter your roll number : ")
    r2 = requests.post (SECURITY_QUESTION_URL, data = {'user_id': roll_no})
    if(r2.text!= 'FALSE'):
        break
    print("Roll number is wrong\nPlease Enter the details again\n")
    time.sleep(1)

password = getpass.getpass("Enter your password : ")
security_ans = getpass.getpass("Answer your  security question - " + r2.text + " : ")

