# Check-for-Go
Seluruh keterangan dari fitur Register Store ini ada di dalam SRS yang tercantum dibawah :
https://docs.google.com/document/d/1v1L9U6XnxSNOSEKYhHNKt-Zhlm9QsXiI/edit?usp=sharing&ouid=116416740980127486591&rtpof=true&sd=true

  int row = jTable2.getSelectedRow();
       
        
        if(row!=-1){
            String cell = jTable2.getModel().getValueAt(row, 0).toString();
            String nomer = jTable2.getValueAt(row, 0).toString();
            String sql = "DELETE FROM stores where toko_id ='"+nomer+"' ";
            
            String resetno = "ALTER TABLE stores DROP toko_id";
            String consecutivenumbers = "ALTER TABLE stores ADD  toko_id NUM( 20 ) NOT NULL AUTO_INCREMENT FIRST ,ADD KEY (toko_id)";
            if (JOptionPane.showConfirmDialog(null,
                            "HAPUS DATA INI..??",
                            "Konfirmasi",JOptionPane.OK_CANCEL_OPTION,
                             JOptionPane.INFORMATION_MESSAGE)
                             == JOptionPane.OK_OPTION){
            try{
           con.createStatement().execute(sql);
           con.createStatement().execute(resetno);
           con.createStatement().execute(consecutivenumbers);
         
           JOptionPane.showMessageDialog(null, "Data Berhasil Di Hapus");
            dispose();
            stores a=new stores();
            a.show();
           
        }catch(Exception e){
          JOptionPane.showMessageDialog(null, e);  }
        }
        }   else{
        JOptionPane.showMessageDialog(null,
                "Klik Salah Satu Tabel");
        
        }