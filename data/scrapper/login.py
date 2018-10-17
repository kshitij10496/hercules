import requests
from bs4 import BeautifulSoup
import getpass
import time


ERP_LOGIN_URL = 'https://erp.iitkgp.ac.in/IIT_ERP3/' #to get the session cookie
SECURITY_QUESTION_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/getSecurityQues.htm'


def get_session_cookie ():
    request_erp = requests.get(ERP_LOGIN_URL)
    
    #Following two lines are needed coz erp has freaking two html tags
    soup = BeautifulSoup(request_erp.text)
    soup = BeautifulSoup(request_erp.text[len(soup):])

    session_tag = soup.find('input',attrs = {'id':'sessionToken'})
    session_cookie = session_tag['value']
    return session_cookie

print(get_session_cookie())

def get_security():
    while (1):
        roll_no = input("Enter your roll number : ")
        r2 = requests.post (SECURITY_QUESTION_URL, data = {'user_id': roll_no})
        if(r2.text!= 'FALSE'):
            break
        print("Roll number is wrong\nPlease Enter the details again\n")
        time.sleep(1)

    password = getpass.getpass("Enter your password : ")
    security_ans = getpass.getpass("Answer your  security question - " + r2.text + " : ")

