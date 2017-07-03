#coding=utf-8
#-*- coding: utf-8 -*-
import MySQLdb as mysqldb
import sys
import gspread

connection = mysqldb.connect('localhost', 'ocsuser', 'ocspassword', 'ocsdb',charset='utf8');
with connection:
    cursor = connection.cursor()
    cursor.execute("select * from accountinfo")
    rows = cursor.fetchall()

from oauth2client.service_account import ServiceAccountCredentials as SAC
GDriveJSON = 'PythonUpload.json'
GSpreadSheet = 'UploadByPython'
try:
    scope = ['https://spreadsheets.google.com/feeds']
    key = SAC.from_json_keyfile_name(GDriveJSON, scope)
    gc = gspread.authorize(key)
    worksheet = gc.open(GSpreadSheet).add_worksheet("accountinfo",1,2)
except Exception as ex:
    print('connect google fail ', ex)
    sys.exit(1)
for row in rows:
    worksheet.append_row((row[0],row[1]))
