import requests
import getpass

ERP_HOME_URL = 'https://erp.iitkgp.ac.in/IIT_ERP3/' #to get the session cookie
SECURITY_QUESTION_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/getSecurityQues.htm'
LOGIN_URL = 'https://erp.iitkgp.ac.in/SSOAdministration/auth.htm'
GET_ACAD_TOKEN_URL = 'https://erp.iitkgp.ac.in/Acad/central_breadth_tt.jsp?action=second'


class Session:
    """
    An erp session

    Logs into erpr after prompting the user to enter details

    
    Class Attributes::
    -self.sessionToken #the erp SessionToken or JSESSIONID (as named by erp)
    -self.ssoToken    #the SSOToken obtained after loggin into erp
    -self.academicToken #the academic token, useful for getting the data from erp

    Basic Usage::

    >>>from erplogin import erpSession
    >>>s = erpSession.Session()
    >>>academicToken = s.academicToken

    """

    headers = {
        'timeout': '20',
        'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/51.0.2704.79 Chrome/51.0.2704.79 Safari/537.36',
    }

    def __init__(self):
        self.sess = requests.Session()
        self.sessionToken = Session.__generate_session_cookie(self)
        self.ssoToken = Session.__generate_sso_token(self)
        self.academicToken = Session.__generate_acad_cookie(self)

    def __generate_session_cookie (self):
        """
        Requests for the session id by requesting for erp home page.

        Arguments:
        > self

        Returns 
        > sessionToken
        """
        print("Getting session cookies")
        response_erp = self.sess.get(ERP_HOME_URL)
        session_token = response_erp.cookies['JSESSIONID']

        return (session_token)


    def __generate_sso_token(self):
        """
        Prompts the user to enter the details and sign into erp with the entered details

        Arguments:
        > self

        Returns 
        > ssoToken
        """
        while (1):
            while (1):
                roll_no = input("Enter your roll number : ")
                response_security_question = self.sess.post (SECURITY_QUESTION_URL, data = {'user_id': roll_no},headers=self.headers)
                if(response_security_question.text!= 'FALSE'):
                    break
                print("Roll number is wrong\nPlease Enter the details again\n")

            password = getpass.getpass("Enter your password : ")
            security_ans = getpass.getpass("Answer your security question - " + response_security_question.text + " : ")

            login_details = {
                            'user_id': roll_no,
                            'password': password,
                            'answer': security_ans,
                            'sessionToken': self.sessionToken,
                            'requestedUrl': 'https://erp.iitkgp.ac.in/IIT_ERP3',
                        }

            response_login_auth_htm = self.sess.post(LOGIN_URL
                                                , data=login_details
                                                , headers=self.headers
                                                )

            if len(response_login_auth_htm.history) >2 :

                ssoToken = response_login_auth_htm.history[1].cookies['ssoToken']
                return ssoToken

            print("Wrong password/security answer\n")

    def __generate_acad_cookie(self):
        """
        Gets the academicToken after loggin into erp
        Arguments:
        > self

        Returns 
        > academicToken
        """
        response_acad = self.sess.get(GET_ACAD_TOKEN_URL)
        academicToken = response_acad.cookies["JSID#/Acad"]
        return academicToken
