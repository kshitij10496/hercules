import requests
import getpass
import time



# Eight cookies have been identified for erp
session_token = ''# 1. JSESSIONID @ path = /IIT_ERP3 : once you open erp.iitkgp.ac.in/IIT_ERP3 it is assigned
                  # 2. JSESSIONID @ path = /SSOAdministration : A useless cookies, just to show how poorly erp-code is written
ssoToken = ''     # 3. ssoToken @ path = / : the cookie that identifies you as the user you logged into your accou
                  # 4. JSESSIONID @ path = /Establishment : Best guess = It is responsible for the aadhar/voter id details pop up 
                  # 5. JSID#/Establishment @ path = /  : Same as the cookie above i.e. 3.
                  # 6. JSID#IITERP3 @ path = / : Same as the cookie 1. mentioned above 
                  # 7. JSESSIONID @ path = /Acad : Useful for the data obtained from erp, pertaining to acads, like time-table, grades.....
academicToken = ''# 8. JSID#ACAD @path = / : Same as above i.e. 7.
# Please contribute to the information above

# *************Cookies are named accordingly*********************
ERP_HOME_URL = 'https://erp.iitkgp.ac.in/IIT_ERP3/' #to get the session cookie
SECURITY_QUESTION_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/getSecurityQues.htm'
#GET_LOGIN_SESSION_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/login.htm?sessionToken={0}&requestedUrl=https://erp.iitkgp.ac.in/IIT_ERP3/home.htm'
LOGIN_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/auth.htm'
GET_ACAD_TOKEN_URL = 'https://erp.iitkgp.ac.in/Acad/central_breadth_tt.jsp?action=second'


headers = {
    'timeout': '20',
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/51.0.2704.79 Chrome/51.0.2704.79 Safari/537.36',
}

s = requests.Session()

def get_session_cookie ():
    print("Getting session cookies")
    response_erp = s.get(ERP_HOME_URL)
    session_token = response_erp.cookies['JSESSIONID']
#    response_login = requests.post(GET_LOGIN_SESSION_URL.format(session_token)
#                               , data = {'sessionToken' : session_token
#                                        ,'requestedUrl' : 'https://erp.iitkgp.ac.in/IIT_ERP3/home.htm'
#                                        }
#                                )
#    cookie2 = response_login.cookies['JSESSIONID']
    return (session_token)


def get_user_details():
    while (1):
        roll_no = input("Enter your roll number : ")
        response_security_question = s.post (SECURITY_QUESTION_URL, data = {'user_id': roll_no},headers=headers)
        if(response_security_question.text!= 'FALSE'):
            break
        print("Roll number is wrong\nPlease Enter the details again\n")
        time.sleep(1)

    password = getpass.getpass("Enter your password : ")
    security_ans = getpass.getpass("Answer your security question - " + response_security_question.text + " : ")

    return roll_no, password, security_ans

def login_into_erp( roll_no, password, security_ans, session_token):
    login_details = {
                    'user_id': roll_no,
                    'password': password,
                    'answer': security_ans,
                    'sessionToken': session_token,
                    'requestedUrl': 'https://erp.iitkgp.ac.in/IIT_ERP3',
                }
    
    response_login_auth_htm = s.post(LOGIN_URL
                                        , data=login_details
                                        , headers=headers
                                        )

    if len(response_login_auth_htm.history) <2 :
        print("Wrong password/security answer\n")
        quit()
    
    sso_token = response_login_auth_htm.history[1].cookies['ssoToken']

#    print(ssoToken)
    return ssoToken

def get_acad_cookie():

    #Get session cookie, by opening the erp page
    session_token = get_session_cookie()
    #Get input from user
    roll_no, password, security_ans = get_user_details()

    #Login to erp with input details and cookies from above
    login_into_erp(roll_no, password, security_ans, session_token)
    
    
    #Get the useful acad session cookie

    response_acad = s.get(GET_ACAD_TOKEN_URL)
    academicToken = response_acad.cookies["JSID#/Acad"]
    return academicToken

def main():
    print("Acad cookie = " + get_acad_cookie())

if __name__ == "__main__":
    main()
