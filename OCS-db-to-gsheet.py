#coding=utf-8
#-*- coding: utf-8 -*-
import MySQLdb as mysqldb
import sys
import gspread
import string

def DB_return(exec_cmd):
    connection = mysqldb.connect('localhost', 'ocsuser', 'ocspass', 'ocsdb',charset='utf8');
    with connection:
        cursor = connection.cursor()
        cursor.execute(exec_cmd)
        rows_info = cursor.fetchall()
        return rows_info

def GSheet_write(GSpreadSheet,GDriveJSON,tab_data,data,worksheet_name):

    a_list=[]
    for (start_row , rowlist) in enumerate(tab_data):
        for (colnum, value) in enumerate(rowlist):
            if colnum == 0:
                a_list.append(str(value))
    a_tup = tuple(a_list)

    start_row = 1
    start_letter = 'A'
    end_len = len(data[0]) - 1
    if end_len > 26:
        len_info = end_len - 26
        len_a='A'
    else:
        len_info = end_len
        len_a=''
    end_letter = string.uppercase[len_info]
    end_letter = len_a+end_letter
    end_row = len(data)
    range = "%s%d:%s%d" % (start_letter, start_row+1 , end_letter, end_row+1)
    range_tab =  "%s%d:%s%d" % (start_letter, start_row , end_letter ,start_row)

    from oauth2client.service_account import ServiceAccountCredentials as SAC
    try:
        scope = ['https://spreadsheets.google.com/feeds']
        key = SAC.from_json_keyfile_name(GDriveJSON, scope)
        gc = gspread.authorize(key)
        wks = gc.open(GSpreadSheet).add_worksheet(worksheet_name,end_row+1,end_len+1)
    except Exception as ex:
        print('connect google fail ', ex)
        sys.exit(1)

    cell_list = wks.range(range)

    try:
        idx = 0
        for (start_row , rowlist) in enumerate(data):
            for (colnum, value) in enumerate(rowlist):
                cell_list[idx].value = value
                idx += 1
                if idx >= len(cell_list):
                    break
        wks.update_cells(cell_list)
    except:
        print "Exception"


    cell_list_tab = wks.range(range_tab)

    try:
        idx = 0
        for (start_row , value) in enumerate(a_tup):
            cell_list_tab[idx].value = value
            idx += 1
            if idx >= len(cell_list_tab):
                break
        wks.update_cells(cell_list_tab)
    except:
        print "Exception"



GDriveJSON = 'PythonUpload.json'
GSpreadSheet = 'UploadByPython'

rows_accountinfo_desc = DB_return("DESCRIBE accountinfo")
rows_accountinfo = DB_return("select * from accountinfo")
GSheet_write(GSpreadSheet,GDriveJSON,rows_accountinfo_desc,rows_accountinfo,"accountinfo")

rows_bios_desc =  DB_return("DESCRIBE bios")
rows_bios = DB_return("select * from bios")
GSheet_write(GSpreadSheet,GDriveJSON,rows_bios_desc,rows_bios,"bios")

rows_cpus_desc =  DB_return("DESCRIBE cpus")
rows_cpus = DB_return("select * from cpus")
GSheet_write(GSpreadSheet,GDriveJSON,rows_cpus_desc,rows_cpus,"cpus")

rows_hardware_desc =  DB_return("DESCRIBE hardware")
rows_hardware = DB_return("select * from hardware")
GSheet_write(GSpreadSheet,GDriveJSON,rows_hardware_desc,rows_hardware,"hardware")
